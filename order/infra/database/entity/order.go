package entity

import (
	"time"
)

type OrderEntity struct {
	ID          int64 `gorm:"autoIncrement;primaryKey"`
	Total       int64 `gorm:"not null"`
	Status      string
	CancelledAt *time.Time
	CompletedAt *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (OrderEntity) TableName() string {
	return "orders"
}
