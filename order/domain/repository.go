package domain

import (
	"context"
)

type Repository interface {
	CreateWithItems(ctx context.Context, order *Order, orderItems []*OrderItem) (*Order, error)
	Update(ctx context.Context, order *Order) error
	GetOrderByID(ctx context.Context, orderID int64) (*Order, error)
}
