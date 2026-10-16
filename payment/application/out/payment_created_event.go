package out

import "time"

type PaymentCreatedEvent struct {
	OrderID           int64      `json:"orderId"`
	Status            string     `json:"status"`
	Amount            int64      `json:"amount"`
	Provider          string     `json:"provider"`
	ProviderReference string     `json:"providerReference"`
	CheckoutURL       string     `json:"checkoutUrl"`
	ExpiresAt         *time.Time `json:"expiresAt,omitempty"`
}
