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

func (o *Order) ChangeStatus(status OrderStatus, now time.Time) error {
	if o.Status == status {
		return nil
	}

	switch status {
	case OrderPending, OrderProcess:
		if o.Status == OrderCancelled || o.Status == OrderSuccess {
			return ErrInvalidOrderStatus
		}
	case OrderCancelled:
		if o.Status == OrderSuccess {
			return ErrInvalidOrderStatus
		}
		o.CancelledAt = &now
	case OrderSuccess:
		if o.Status == OrderCancelled {
			return ErrInvalidOrderStatus
		}
		o.CompletedAt = &now
	default:
		return ErrInvalidOrderStatus
	}

	o.Status = status
	o.UpdatedAt = now

	return nil
}

func (o *Order) Completed() error {
	return o.ChangeStatus(OrderSuccess, time.Now().UTC())
}

func (o *Order) Cancel() error {
	return o.ChangeStatus(OrderCancelled, time.Now().UTC())
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
