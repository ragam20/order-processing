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
		ID:                    payment.ID,
		OrderID:               orderID,
		Amount:                payment.Amount,
		Status:                payment.Status,
		Provider:              payment.Provider,
		ProviderReference:     payment.ProviderReference,
		ProviderTransactionID: payment.ProviderTransactionID,
		CheckoutURL:           payment.CheckoutURL,
		PaymentToken:          payment.PaymentToken,
		ResponseProvider:      payment.ResponseProvider,
		PaidAt:                payment.PaidAt,
		CreatedAt:             payment.CreatedAt,
		UpdatedAt:             payment.UpdatedAt,
	}
	if err := r.db.WithContext(ctx).Create(&e).Error; err != nil {
		return err
	}

	return nil
}

func (r *repository) GetByProviderReference(ctx context.Context, providerReference string) (*domain.Payment, error) {
	var e entity.PaymentEntity
	if err := r.db.WithContext(ctx).Where("provider_reference = ?", providerReference).First(&e).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, domain.ErrPaymentNotFound
		}
		return nil, err
	}

	return &domain.Payment{
		ID:                    e.ID,
		OrderID:               e.OrderID,
		Amount:                e.Amount,
		Status:                e.Status,
		Provider:              e.Provider,
		ProviderReference:     e.ProviderReference,
		ProviderTransactionID: e.ProviderTransactionID,
		CheckoutURL:           e.CheckoutURL,
		PaymentToken:          e.PaymentToken,
		ResponseProvider:      e.ResponseProvider,
		PaidAt:                e.PaidAt,
		CreatedAt:             e.CreatedAt,
		UpdatedAt:             e.UpdatedAt,
	}, nil
}

func (r *repository) Update(ctx context.Context, payment *domain.Payment) error {
	e := entity.PaymentEntity{
		ID:                    payment.ID,
		OrderID:               payment.OrderID,
		Amount:                payment.Amount,
		Status:                payment.Status,
		Provider:              payment.Provider,
		ProviderReference:     payment.ProviderReference,
		ProviderTransactionID: payment.ProviderTransactionID,
		CheckoutURL:           payment.CheckoutURL,
		PaymentToken:          payment.PaymentToken,
		ResponseProvider:      payment.ResponseProvider,
		PaidAt:                payment.PaidAt,
		CreatedAt:             payment.CreatedAt,
		UpdatedAt:             payment.UpdatedAt,
	}
	if err := r.db.WithContext(ctx).Model(&e).Updates(&e).Error; err != nil {
		return err
	}

	return nil
}
