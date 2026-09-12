package domain

import (
	"time"
)

type OrderStatus string

var (
	OrderPending   OrderStatus = "pending"
	OrderProcess   OrderStatus = "process"
	OrderCancelled OrderStatus = "cancelled"
	OrderSuccess   OrderStatus = "success"
)

type Order struct {
	ID          int64
	Total       int64
	Status      OrderStatus
	CancelledAt *time.Time
	CompletedAt *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (o *Order) Completed() error {
	if o.Status == OrderSuccess {
		return nil
	}

	if o.Status != OrderCancelled {
		return ErrOrderSuccess
	}

	t := time.Now()
	o.CompletedAt = &t
	o.Status = OrderSuccess

	return nil
}

func (o *Order) Cancel() error {
	if o.Status != OrderProcess {
		return ErrOrderCanceled
	}

	t := time.Now()
	o.CancelledAt = &t
	o.Status = OrderCancelled

	return nil
}

type OrderItem struct {
	ID        int64
	ProductID int64

	Name     string
	Price    int64
	Quantity int32

	CreatedAt time.Time
	UpdatedAt time.Time
}
