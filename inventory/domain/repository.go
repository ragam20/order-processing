package domain

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, product *Product) error
	Update(ctx context.Context, product *Product) error
	GetProducts(ctx context.Context) []Product
	GetProductByID(ctx context.Context, productID int64) (*Product, error)
}
