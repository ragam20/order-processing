package domain

import "context"

type Repository interface {
	Create(ctx context.Context, payment *Payment, orderID int64) error
	Update(ctx context.Context, payment *Payment, orderID int64) error
}
