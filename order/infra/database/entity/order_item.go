package entity

import (
	"time"

	"github.com/shopspring/decimal"
)

type OrderItemEntity struct {
	ID int64 `gorm:"autoIncrement;primaryKey"`

	OrderID int64

	Name     string
	Price    decimal.Decimal `gorm:"type:decimal(20,2);not null"`
	Quantity int32

	CreatedAt time.Time
	UpdatedAt time.Time

	Order OrderEntity `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:OrderID;reference:ID"`
}

func (OrderItemEntity) TableName() string {
	return "order_items"
}
