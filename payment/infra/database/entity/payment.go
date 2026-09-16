package entity

import (
	"encoding/json"
	"time"

	"ops/payment/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PaymentEntity struct {
	ID      uuid.UUID `gorm:"type:uuid;primaryKey"`
	OrderID int64     `gorm:"not null"`
	Amount  int64     `gorm:"not null"`
	Method  string
	Status  domain.PaymentStatus `gorm:"not null;type:varchar(25);check:payment_status_check,status IN ('AUTHORIZED', 'PENDING', 'PAID', 'PROCESSING', 'FAILED', 'CANCELLED', 'EXPIRED', 'REFUND')"`

	ProviderReference     string `gorm:"index"`
	ProviderTransactionID string `gorm:"index"`
	Provider              string `gorm:"not null"`

	CheckoutURL      string `gorm:"not null"`
	PaymentToken     string
	ResponseProvider json.RawMessage
	PaidAt           *time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (PaymentEntity) TableName() string {
	return "payments"
}

func (p *PaymentEntity) BeforeCreate(tx *gorm.DB) error {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}

	return nil
}
