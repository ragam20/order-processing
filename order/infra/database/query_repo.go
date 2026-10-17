package database

import (
	"context"

	"ops/order/application/out"
	"ops/order/domain"
	"ops/order/infra/database/entity"

	"gorm.io/gorm"
)

func (r *Repository) GetOrdersWithItems(ctx context.Context) ([]out.OrderWithItemsResult, error) {
	return r.getOrdersWithItems(ctx, r.db.WithContext(ctx).Order("id DESC"))
}

func (r *Repository) GetOrderDetail(ctx context.Context, orderID int64) (*out.OrderWithItemsResult, error) {
	orders, err := r.getOrdersWithItems(ctx, r.db.WithContext(ctx).Where("id = ?", orderID))
	if err != nil {
		return nil, err
	}
	if len(orders) == 0 {
		return nil, domain.ErrOrderNotFound
	}

	return &orders[0], nil
}

func (r *Repository) GetOrdersByStatus(ctx context.Context, status domain.OrderStatus) ([]out.OrderWithItemsResult, error) {
	return r.getOrdersWithItems(
		ctx,
		r.db.WithContext(ctx).
			Where("status = ?", string(status)).
			Order("id DESC"),
	)
}

func (r *Repository) getOrdersWithItems(ctx context.Context, query *gorm.DB) ([]out.OrderWithItemsResult, error) {
	var orderEntities []entity.OrderEntity
	if err := query.Find(&orderEntities).Error; err != nil {
		return []out.OrderWithItemsResult{}, err
	}

	result := make([]out.OrderWithItemsResult, 0, len(orderEntities))
	if len(orderEntities) == 0 {
		return result, nil
	}

	orderIDs := make([]int64, 0, len(orderEntities))
	for _, orderEntity := range orderEntities {
		orderIDs = append(orderIDs, orderEntity.ID)
	}

	var itemEntities []entity.OrderItemEntity
	if err := r.db.WithContext(ctx).
		Where("order_id IN ?", orderIDs).
		Order("id ASC").
		Find(&itemEntities).Error; err != nil {
		return []out.OrderWithItemsResult{}, err
	}

	itemsByOrderID := make(map[int64][]domain.OrderItem, len(orderIDs))
	for _, itemEntity := range itemEntities {
		itemsByOrderID[itemEntity.OrderID] = append(itemsByOrderID[itemEntity.OrderID], domain.OrderItem{
			ID:        itemEntity.ID,
			ProductID: itemEntity.ProductID,
			Name:      itemEntity.Name,
			Price:     itemEntity.Price,
			Quantity:  itemEntity.Quantity,
			CreatedAt: itemEntity.CreatedAt,
			UpdatedAt: itemEntity.UpdatedAt,
		})
	}

	for _, orderEntity := range orderEntities {
		items := itemsByOrderID[orderEntity.ID]
		if items == nil {
			items = []domain.OrderItem{}
		}

		result = append(result, out.OrderWithItemsResult{
			Order: domain.Order{
				ID:          orderEntity.ID,
				Total:       orderEntity.Total,
				Status:      domain.OrderStatus(orderEntity.Status),
				CancelledAt: orderEntity.CancelledAt,
				CompletedAt: orderEntity.CompletedAt,
				CreatedAt:   orderEntity.CreatedAt,
				UpdatedAt:   orderEntity.UpdatedAt,
			},
			Items: items,
		})
	}

	return result, nil
}
