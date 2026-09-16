package out

type PaymentStatusChangedEvent struct {
	OrderID               int64  `json:"orderId"`
	Status                string `json:"status"`
	Amount                int64  `json:"amount"`
	Provider              string `json:"provider"`
	ProviderReference     string `json:"providerReference"`
	ProviderTransactionID string `json:"providerTransactionId"`
}
