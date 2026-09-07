package entity

import (
	"time"

	"github.com/shopspring/decimal"
)

type OrderEntity struct {
	ID          int64           `gorm:"autoIncrement;primaryKey"`
	Total       decimal.Decimal `gorm:"type:decimal(20,2);not null"`
	Status      string
	CancelledAt *time.Time
	CompletedAt *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (OrderEntity) TableName() string {
	return "orders"
}
