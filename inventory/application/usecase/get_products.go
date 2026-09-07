package usecase

import (
	"context"

	"ops/inventory/domain"
)

type GetProductsUseCase struct {
	repo domain.Repository
}

func NewGetProductsUseCase(repo domain.Repository) *GetProductsUseCase {
	return &GetProductsUseCase{repo: repo}
}

func (u *GetProductsUseCase) Execute(ctx context.Context) []domain.Product {
	return u.repo.GetProducts(ctx)
}
