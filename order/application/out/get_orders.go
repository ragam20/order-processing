package out

import "ops/order/domain"

type OrderWithItemsResult struct {
	Order domain.Order
	Items []domain.OrderItem
}
