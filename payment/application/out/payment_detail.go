package out

import (
	"time"

	"github.com/google/uuid"
)

type PaymentDetailResult struct {
	ID                    uuid.UUID
	OrderID               int64
	Amount                int64
	Status                string
	Provider              string
	ProviderReference     string
	ProviderTransactionID string
	CheckoutURL           string
	PaymentToken          string
	PaidAt                *time.Time
	CreatedAt             time.Time
	UpdatedAt             time.Time
}
