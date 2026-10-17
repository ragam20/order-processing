package response

import (
	"time"

	"github.com/google/uuid"
)

type CreatePaymentResponse struct {
	Status                string `json:"status"`
	ProviderReference     string `json:"providerReference"`
	ProviderTransactionID string `json:"providerTransactionId"`
	CheckoutURL           string `json:"checkoutUrl"`
	PaymentToken          string `json:"paymentToken"`
}

type GetStatusPaymentResponse struct {
	Status                string     `json:"status"`
	Provider              string     `json:"provider"`
	ProviderReference     string     `json:"providerReference"`
	ProviderTransactionID string     `json:"providerTransactionId"`
	Amount                int64      `json:"amount"`
	PaidAt                *time.Time `json:"paidAt"`
}

type GetPaymentByOrderIDResponse struct {
	ID                    uuid.UUID  `json:"id"`
	OrderID               int64      `json:"orderId"`
	Amount                int64      `json:"amount"`
	Status                string     `json:"status"`
	Provider              string     `json:"provider"`
	ProviderReference     string     `json:"providerReference"`
	ProviderTransactionID string     `json:"providerTransactionId"`
	CheckoutURL           string     `json:"checkoutUrl"`
	PaymentToken          string     `json:"paymentToken"`
	PaidAt                *time.Time `json:"paidAt"`
	CreatedAt             time.Time  `json:"createdAt"`
	UpdatedAt             time.Time  `json:"updatedAt"`
}
