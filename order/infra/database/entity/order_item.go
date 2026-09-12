package entity

import (
	"time"
)

type OrderItemEntity struct {
	ID int64 `gorm:"autoIncrement;primaryKey"`

	OrderID   int64
	ProductID int64

	Name     string
	Price    int64 `gorm:"not null"`
	Quantity int32

	CreatedAt time.Time
	UpdatedAt time.Time

	Order OrderEntity `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:OrderID;reference:ID"`
}

func (OrderItemEntity) TableName() string {
	return "order_items"
}
