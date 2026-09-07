package out

import (
	"encoding/json"
	"time"

	"ops/payment/domain"
)

type PaymentStatusResult struct {
	ProviderReference     string
	ProviderTransactionID string
	Status                domain.PaymentStatus
	ProviderStatus        string
	GrossAmount           int64
	PaidAt                *time.Time
	RawResponse           json.RawMessage
}
