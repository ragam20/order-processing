package usecase

import (
	"context"

	"ops/order/application/in"
	"ops/order/domain"
)

type CheckoutUseCase struct {
	repo domain.Repository
}

func NewCheckoutUseCase(repo domain.Repository) *CheckoutUseCase {
	return &CheckoutUseCase{repo: repo}
}

func (u *CheckoutUseCase) Execute(ctx context.Context, input in.CheckoutInput) error {
	return nil
}
