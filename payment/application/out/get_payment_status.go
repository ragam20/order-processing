package out

import "time"

type GetStatusPaymentResult struct {
	Status                string
	ProviderReference     string
	Provider              string
	ProviderTransactionID string
	Amount                int64
	PaidAt                *time.Time
}
