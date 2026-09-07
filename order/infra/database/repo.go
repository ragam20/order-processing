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

func (r *repository) Create(ctx context.Context, order *domain.Order) error {
	e := entity.OrderEntity{
		ID:        order.ID,
		Total:     order.Total,
		CreatedAt: order.CreatedAt,
		UpdatedAt: order.UpdatedAt,
	}

	if err := r.db.WithContext(ctx).
		Create(&e).
		Error; err != nil {
		return err
	}

	order.ID = e.ID

	return nil
}

func (r *repository) Update(ctx context.Context, order *domain.Order) error {
	e := entity.OrderEntity{
		ID:        order.ID,
		Total:     order.Total,
		UpdatedAt: order.UpdatedAt,
	}

	return r.db.WithContext(ctx).
		Model(&e).
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
		ID:        e.ID,
		Total:     e.Total,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
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
			ID:        en.ID,
			Total:     en.Total,
			CreatedAt: en.CreatedAt,
			UpdatedAt: en.UpdatedAt,
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
			Name:      en.Name,
			Price:     en.Price,
			Quantity:  en.Quantity,
			CreatedAt: en.CreatedAt,
			UpdatedAt: en.UpdatedAt,
		})
	}

	return d
}
