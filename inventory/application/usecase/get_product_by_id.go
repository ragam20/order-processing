package usecase

import (
	"context"

	"ops/inventory/domain"
)

type GetProductByIDUseCase struct {
	repo domain.Repository
}

func NewGetProductByIDUseCase(repo domain.Repository) *GetProductByIDUseCase {
	return &GetProductByIDUseCase{repo: repo}
}

func (u *GetProductByIDUseCase) Execute(ctx context.Context, productID int64) (*domain.Product, error) {
	p, err := u.repo.GetProductByID(ctx, productID)
	if err != nil {
		return nil, err
	}

	return p, nil
}
