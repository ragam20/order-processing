package domain

import "context"

type Repository interface {
	Create(ctx context.Context, order *Order) error
	Update(ctx context.Context, order *Order) error
	GetOrderByID(ctx context.Context, orderID int64) (*Order, error)
	GetOrders(ctx context.Context) []Order
	GetItems(ctx context.Context, orderID int64) []OrderItem
}
