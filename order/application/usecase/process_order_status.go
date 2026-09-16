package usecase

import (
	"context"

	"ops/order/domain"
)

type ProcessOrderStatusUseCase struct {
	repo domain.Repository
}

func NewProcessOrderStatusUseCase(repo domain.Repository) *ProcessOrderStatusUseCase {
	return &ProcessOrderStatusUseCase{repo: repo}
}

func (u *ProcessOrderStatusUseCase) Execute(ctx context.Context) error {
	return nil
}
