package usecase

import (
	"context"

	"ops/payment/domain"
)

type CreatePaymentUseCase struct {
	repo domain.Repository
}

func NewCreatePaymentUseCase(repo domain.Repository) *CreatePaymentUseCase {
	return &CreatePaymentUseCase{repo: repo}
}

func (u *CreatePaymentUseCase) Execute(ctx context.Context) error {
	return nil
}
