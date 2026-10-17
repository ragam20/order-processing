package database

import (
	"context"
	"errors"

	"ops/order/domain"
	"ops/order/infra/database/entity"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateWithItems(ctx context.Context, order *domain.Order, orderItems []*domain.OrderItem) (*domain.Order, error) {
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		orderEntity := entity.OrderEntity{
			Total:     order.Total,
			Status:    string(order.Status),
			CreatedAt: order.CreatedAt,
			UpdatedAt: order.UpdatedAt,
		}

		if err := tx.Create(&orderEntity).Error; err != nil {
			return err
		}

		order.ID = orderEntity.ID

		itemEntities := make([]*entity.OrderItemEntity, 0, len(orderItems))
		for _, item := range orderItems {
			itemEntities = append(itemEntities, &entity.OrderItemEntity{
				OrderID:   order.ID,
				ProductID: item.ProductID,
				Name:      item.Name,
				Quantity:  item.Quantity,
				Price:     item.Price,
			})
		}

		return tx.Create(&itemEntities).Error
	})
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (r *Repository) Update(ctx context.Context, order *domain.Order) error {
	e := entity.OrderEntity{
		ID:          order.ID,
		Total:       order.Total,
		Status:      string(order.Status),
		CancelledAt: order.CancelledAt,
		CompletedAt: order.CompletedAt,
		UpdatedAt:   order.UpdatedAt,
	}

	return r.db.WithContext(ctx).
		Model(&entity.OrderEntity{}).
		Where("id = ?", order.ID).
		Updates(&e).
		Error
}

func (r *Repository) GetOrderByID(ctx context.Context, orderID int64) (*domain.Order, error) {
	var e entity.OrderEntity
	if err := r.db.WithContext(ctx).Where("id = ?", orderID).First(&e).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrOrderNotFound
		}
		return nil, err
	}

	return &domain.Order{
		ID:          e.ID,
		Total:       e.Total,
		Status:      domain.OrderStatus(e.Status),
		CancelledAt: e.CancelledAt,
		CompletedAt: e.CompletedAt,
		CreatedAt:   e.CreatedAt,
		UpdatedAt:   e.UpdatedAt,
	}, nil
}
