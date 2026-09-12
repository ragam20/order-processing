package database

import (
	"context"
	"errors"

	"ops/order/domain"
	"ops/order/infra/database/entity"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) domain.Repository {
	return &repository{db: db}
}

func (r *repository) CreateWithItems(ctx context.Context, order *domain.Order, orderItems []*domain.OrderItem) (*domain.Order, error) {
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

func (r *repository) Update(ctx context.Context, order *domain.Order) error {
	e := entity.OrderEntity{
		ID:        order.ID,
		Total:     order.Total,
		Status:    string(order.Status),
		UpdatedAt: order.UpdatedAt,
	}

	return r.db.WithContext(ctx).
		Model(&entity.OrderEntity{}).
		Where("id = ?", order.ID).
		Updates(&e).
		Error
}

func (r *repository) GetOrderByID(ctx context.Context, orderID int64) (*domain.Order, error) {
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

func (r *repository) GetOrders(ctx context.Context) []domain.Order {
	var e []entity.OrderEntity
	if err := r.db.WithContext(ctx).Find(&e).Error; err != nil {
		return []domain.Order{}
	}

	d := make([]domain.Order, 0)
	for _, en := range e {
		d = append(d, domain.Order{
			ID:          en.ID,
			Total:       en.Total,
			Status:      domain.OrderStatus(en.Status),
			CancelledAt: en.CancelledAt,
			CompletedAt: en.CompletedAt,
			CreatedAt:   en.CreatedAt,
			UpdatedAt:   en.UpdatedAt,
		})
	}

	return d
}

func (r *repository) GetItems(ctx context.Context, orderID int64) []domain.OrderItem {
	var e []entity.OrderItemEntity
	if err := r.db.WithContext(ctx).Where("order_id = ?", orderID).Find(&e).Error; err != nil {
		return []domain.OrderItem{}
	}

	d := make([]domain.OrderItem, 0)
	for _, en := range e {
		d = append(d, domain.OrderItem{
			ID:        en.ID,
			ProductID: en.ProductID,
			Name:      en.Name,
			Price:     en.Price,
			Quantity:  en.Quantity,
			CreatedAt: en.CreatedAt,
			UpdatedAt: en.UpdatedAt,
		})
	}

	return d
}
