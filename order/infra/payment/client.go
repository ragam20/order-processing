package payment

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"ops/order/application"
	"ops/order/application/in"
	"ops/order/application/out"
	"ops/shared/config"
)

type paymentClient struct {
	env *config.Env
}

func New(env *config.Env) application.PaymentService {
	return &paymentClient{env: env}
}

func (c *paymentClient) CreatePayment(ctx context.Context, input in.CreatePaymentInput) (out.PaymentResult, error) {
	items := make([]PaymentItemRequest, 0, len(input.Items))
	for _, item := range input.Items {
		items = append(items, PaymentItemRequest{
			ID: item.ID, Name: item.Name, Price: item.Price, Quantity: item.Quantity,
		})
	}

	body, err := json.Marshal(CreatePaymentRequest{
		OrderID: input.OrderID, TotalAmount: input.TotalAmount, Items: items,
		Customer: CustomerRequest{
			ID: input.Customer.ID, Name: input.Customer.Name,
			Email: input.Customer.Email, Phone: input.Customer.Phone,
		},
		FinishURL: input.FinishURL,
	})
	if err != nil {
		return out.PaymentResult{}, err
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		c.env.HTTP.PaymentURL+"/api/internal/payment",
		bytes.NewReader(body),
	)
	if err != nil {
		return out.PaymentResult{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return out.PaymentResult{}, err
	}
	defer resp.Body.Close()

	var result CreatePaymentResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return out.PaymentResult{}, err
	}
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices || !result.Success {
		return out.PaymentResult{}, fmt.Errorf("payment service returned status: %d", resp.StatusCode)
	}

	var token *string
	if result.Data.PaymentToken != "" {
		value := result.Data.PaymentToken
		token = &value
	}

	return out.PaymentResult{
		Status:                result.Data.Status,
		ProviderReference:     result.Data.ProviderReference,
		ProviderTransactionID: result.Data.ProviderTransactionID,
		CheckoutURL:           result.Data.CheckoutURL,
		PaymentToken:          token,
	}, nil
}

func (c *paymentClient) GetStatusPayment(ctx context.Context, providerReference string) (out.GetStatusPaymentResult, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		c.env.HTTP.PaymentURL+"/api/internal/payment/status/"+url.PathEscape(providerReference),
		nil,
	)
	if err != nil {
		return out.GetStatusPaymentResult{}, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return out.GetStatusPaymentResult{}, err
	}
	defer resp.Body.Close()

	var result GetStatusPaymentResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return out.GetStatusPaymentResult{}, err
	}
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices || !result.Success {
		return out.GetStatusPaymentResult{}, fmt.Errorf("payment service returned status: %d", resp.StatusCode)
	}

	return out.GetStatusPaymentResult{
		Status:                result.Data.Status,
		Provider:              result.Data.Provider,
		ProviderReference:     result.Data.ProviderReference,
		ProviderTransactionID: result.Data.ProviderTransactionID,
		Amount:                result.Data.Amount,
		PaidAt:                result.Data.PaidAt,
	}, nil
}
