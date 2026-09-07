package usecase

import "context"

type ProcessPaymentUseCase struct{}

func NewProcessPaymentUseCase() *ProcessPaymentUseCase {
	return &ProcessPaymentUseCase{}
}

func (u *ProcessPaymentUseCase) Execute(ctx context.Context) {}
