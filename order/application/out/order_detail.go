package out

import "ops/order/domain"

type OrderDetailResult struct {
	Order *domain.Order
	Items []domain.OrderItem
}
