package in

import (
	"encoding/json"
	"time"

	"ops/payment/domain"
)

type ProcessPaymentInput struct {
	Provider              string
	ProviderReference     string
	ProviderTransactionID string
	Status                domain.PaymentStatus
	ProviderStatus        string
	GrossAmount           int64
	PaidAt                *time.Time
	RawPayload            json.RawMessage
}
