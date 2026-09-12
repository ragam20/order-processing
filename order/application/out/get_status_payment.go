package out

import "time"

type GetStatusPaymentResult struct {
	OrderID               int64
	Status                string
	Provider              string
	ProviderReference     string
	ProviderTransactionID string
	Amount                int64
	PaidAt                *time.Time
}
