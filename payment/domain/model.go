package domain

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type PaymentStatus string

var (
	PaymentStatusAuthorized PaymentStatus = "authorized"
	PaymentStatusPending    PaymentStatus = "pending"
	PaymentStatusPaid       PaymentStatus = "paid"
	PaymentStatusProcessing PaymentStatus = "processing"
	PaymentStatusFailed     PaymentStatus = "failed"
	PaymentStatusCancelled  PaymentStatus = "cancelled"
	PaymentStatusExpired    PaymentStatus = "expired"
	PaymentStatusRefund     PaymentStatus = "refund"
)

type Payment struct {
	ID                    uuid.UUID
	Status                PaymentStatus
	Provider              string
	ProviderTransactionID string

	CheckoutURL      string
	ResponseProvider json.RawMessage

	CreatedAt time.Time
	UpdatedAt time.Time
}
