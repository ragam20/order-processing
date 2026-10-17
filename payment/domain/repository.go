package domain

import "context"

type Repository interface {
	Create(ctx context.Context, payment *Payment, orderID int64) error
	GetByOrderID(ctx context.Context, orderID int64) (*Payment, error)
	GetByProviderReference(ctx context.Context, providerReference string) (*Payment, error)
	Update(ctx context.Context, payment *Payment) error
}
