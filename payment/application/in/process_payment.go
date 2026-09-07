package in

import (
	"encoding/json"

	"ops/payment/domain"
)

type ProcessPaymentInput struct {
	Provider      string
	OrderID       string
	TransactionID string
	Status        domain.PaymentStatus
	RawPayload    json.RawMessage
}
