package database

import (
	"context"

	"ops/payment/domain"
	"ops/payment/infra/database/entity"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) domain.Repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, payment *domain.Payment, orderID int64) error {
	e := entity.PaymentEntity{
		ID:               payment.ID,
		OrderID:          orderID,
		Status:           payment.Status,
		Provider:         payment.Provider,
		ResponseProvider: payment.ResponseProvider,
		CreatedAt:        payment.CreatedAt,
		UpdatedAt:        payment.UpdatedAt,
	}
	if err := r.db.WithContext(ctx).Create(&e).Error; err != nil {
		return err
	}

	return nil
}

func (r *repository) Update(ctx context.Context, payment *domain.Payment, orderID int64) error {
	e := entity.PaymentEntity{
		ID:               payment.ID,
		OrderID:          orderID,
		Status:           payment.Status,
		Provider:         payment.Provider,
		ResponseProvider: payment.ResponseProvider,
		CreatedAt:        payment.CreatedAt,
		UpdatedAt:        payment.UpdatedAt,
	}
	if err := r.db.WithContext(ctx).Model(&e).Updates(&e).Error; err != nil {
		return err
	}

	return nil
}
