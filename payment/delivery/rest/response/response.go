package response

import "time"

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
