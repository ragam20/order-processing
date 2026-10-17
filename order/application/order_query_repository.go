package application

import (
	"context"

	"ops/order/application/out"
	"ops/order/domain"
)

type OrderQueryRepository interface {
	GetOrdersWithItems(ctx context.Context) ([]out.OrderWithItemsResult, error)
	GetOrderDetail(ctx context.Context, orderID int64) (*out.OrderWithItemsResult, error)
	GetOrdersByStatus(ctx context.Context, status domain.OrderStatus) ([]out.OrderWithItemsResult, error)
}
