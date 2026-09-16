package domain

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type PaymentStatus string

var (
	PaymentStatusAuthorized PaymentStatus = "AUTHORIZED"
	PaymentStatusPending    PaymentStatus = "PENDING"
	PaymentStatusPaid       PaymentStatus = "PAID"
	PaymentStatusProcessing PaymentStatus = "PROCESSING"
	PaymentStatusFailed     PaymentStatus = "FAILED"
	PaymentStatusCancelled  PaymentStatus = "CANCELLED"
	PaymentStatusExpired    PaymentStatus = "EXPIRED"
	PaymentStatusRefund     PaymentStatus = "REFUND"
)

type Payment struct {
	ID                    uuid.UUID
	OrderID               int64
	Amount                int64
	Status                PaymentStatus
	Provider              string
	ProviderReference     string
	ProviderTransactionID string

	CheckoutURL      string
	PaymentToken     string
	ResponseProvider json.RawMessage
	PaidAt           *time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
}
