package midtrans

import (
	"context"
	"crypto/sha512"
	"crypto/subtle"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"ops/payment/application"
	"ops/payment/application/in"
	"ops/payment/application/out"
	"ops/payment/domain"
	"ops/shared/config"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"
	"github.com/shopspring/decimal"
)

type midtransProvider struct {
	client    snap.Client
	core      coreapi.Client
	serverKey string
}

func New(env *config.Env) application.PaymentProvider {
	var snapClient snap.Client
	snapClient.New(env.Midtrans.ServerKey, midtrans.Sandbox)

	var coreClient coreapi.Client
	coreClient.New(env.Midtrans.ServerKey, midtrans.Sandbox)

	return &midtransProvider{
		client:    snapClient,
		core:      coreClient,
		serverKey: env.Midtrans.ServerKey,
	}
}

func (mt *midtransProvider) GetProviderName() string {
	return "Midtrans"
}

func (mt *midtransProvider) CreatePayment(ctx context.Context, input in.CreatePaymentInput) (out.PaymentResult, error) {
	var mdItems *[]midtrans.ItemDetails
	result := make([]midtrans.ItemDetails, 0, len(input.Items))

	for _, item := range input.Items {
		result = append(result, midtrans.ItemDetails{
			ID:    item.ID,
			Name:  item.Name,
			Price: item.Price,
			Qty:   item.Quantity,
		})
	}

	mdItems = &result

	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.FormatInt(input.OrderID, 10),
			GrossAmt: input.TotalAmount,
		},
		Items: mdItems,
		CustomerDetail: &midtrans.CustomerDetails{
			FName: input.Customer.Name,
			Email: input.Customer.Email,
			Phone: input.Customer.Phone,
		},
	}

	if input.FinishURL != "" {
		req.Callbacks = &snap.Callbacks{Finish: input.FinishURL}
	}

	resp, err := mt.client.CreateTransaction(req)
	if err != nil {
		return out.PaymentResult{}, err
	}

	rawResponse, marshalErr := json.Marshal(resp)
	if marshalErr != nil {
		return out.PaymentResult{}, fmt.Errorf("marshal midtrans response: %w", marshalErr)
	}

	return out.PaymentResult{
		ProviderReference: strconv.FormatInt(input.OrderID, 10),
		CheckoutURL:       resp.RedirectURL,
		PaymentToken:      resp.Token,
		RawResponse:       rawResponse,
	}, nil
}

func (mt *midtransProvider) GetStatus(ctx context.Context, reference string) (out.PaymentStatusResult, error) {
	resp, providerErr := mt.core.CheckTransaction(reference)
	if providerErr != nil {
		return out.PaymentStatusResult{}, providerErr
	}

	status, statusErr := mapStatus(resp.TransactionStatus, resp.FraudStatus)
	if statusErr != nil {
		return out.PaymentStatusResult{}, statusErr
	}

	grossAmount, amountErr := parseGrossAmount(resp.GrossAmount)
	if amountErr != nil {
		return out.PaymentStatusResult{}, amountErr
	}

	rawResponse, marshalErr := json.Marshal(resp)
	if marshalErr != nil {
		return out.PaymentStatusResult{}, fmt.Errorf("marshal midtrans status response: %w", marshalErr)
	}

	return out.PaymentStatusResult{
		ProviderReference:     resp.OrderID,
		ProviderTransactionID: resp.TransactionID,
		Status:                status,
		ProviderStatus:        resp.TransactionStatus,
		GrossAmount:           grossAmount,
		PaidAt:                paymentTime(resp.SettlementTime),
		RawResponse:           rawResponse,
	}, nil
}

func (mt *midtransProvider) ParseWebhook(input in.ParseWebhook) (out.PaymentNotificationResult, error) {
	var notification webhookNotification
	if err := json.Unmarshal(input.Payload, &notification); err != nil {
		return out.PaymentNotificationResult{}, fmt.Errorf("decode midtrans webhook: %w", err)
	}

	if !mt.validSignature(notification) {
		return out.PaymentNotificationResult{}, fmt.Errorf("invalid midtrans webhook signature")
	}

	status, err := mapStatus(notification.TransactionStatus, notification.FraudStatus)
	if err != nil {
		return out.PaymentNotificationResult{}, err
	}

	grossAmount, err := parseGrossAmount(notification.GrossAmount)
	if err != nil {
		return out.PaymentNotificationResult{}, err
	}

	return out.PaymentNotificationResult{
		ProviderReference:     notification.OrderID,
		ProviderTransactionID: notification.TransactionID,
		Status:                status,
		ProviderStatus:        notification.TransactionStatus,
		GrossAmount:           grossAmount,
		PaidAt:                paymentTime(notification.SettlementTime),
		RawPayload:            json.RawMessage(input.Payload),
	}, nil
}

type webhookNotification struct {
	OrderID           string `json:"order_id"`
	StatusCode        string `json:"status_code"`
	GrossAmount       string `json:"gross_amount"`
	SignatureKey      string `json:"signature_key"`
	TransactionID     string `json:"transaction_id"`
	TransactionStatus string `json:"transaction_status"`
	FraudStatus       string `json:"fraud_status"`
	SettlementTime    string `json:"settlement_time"`
}

func (mt *midtransProvider) validSignature(notification webhookNotification) bool {
	value := notification.OrderID + notification.StatusCode + notification.GrossAmount + mt.serverKey
	hash := sha512.Sum512([]byte(value))
	expected := hex.EncodeToString(hash[:])
	actual := strings.ToLower(notification.SignatureKey)

	return subtle.ConstantTimeCompare([]byte(expected), []byte(actual)) == 1
}

// Mapping transcation status lifecycle Midtrans into our backend lifecycle
func mapStatus(transactionStatus string, fraudStatus string) (domain.PaymentStatus, error) {
	switch strings.ToLower(transactionStatus) {
	case "authorize":
		return domain.PaymentStatusAuthorized, nil
	case "settlement":
		return domain.PaymentStatusPaid, nil
	case "capture":
		if strings.EqualFold(fraudStatus, "accept") {
			return domain.PaymentStatusPaid, nil
		}
		return domain.PaymentStatusPending, nil
	case "pending":
		return domain.PaymentStatusPending, nil
	case "expire":
		return domain.PaymentStatusExpired, nil
	case "deny", "failure":
		return domain.PaymentStatusFailed, nil
	case "cancel":
		return domain.PaymentStatusCancelled, nil
	case "refund", "partial_refund":
		return domain.PaymentStatusRefund, nil
	default:
		return "", fmt.Errorf("unsupported midtrans transaction status: %q", transactionStatus)
	}
}

func parseGrossAmount(value string) (int64, error) {
	amount, err := decimal.NewFromString(value)
	if err != nil {
		return 0, fmt.Errorf("parse midtrans gross amount %q: %w", value, err)
	}

	return amount.IntPart(), nil
}

func paymentTime(value string) *time.Time {
	if value == "" {
		return nil
	}

	for _, layout := range []string{
		"2006-01-02 15:04:05",
		time.RFC3339,
	} {
		parsed, err := time.Parse(layout, value)
		if err == nil {
			return &parsed
		}
	}

	return nil
}
