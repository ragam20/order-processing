package usecase

import (
	"context"
	"time"

	"ops/order/application/in"
	"ops/order/domain"
	"ops/shared/app"
)

type ProcessOrderStatusUseCase struct {
	repo domain.Repository
}

func NewProcessOrderStatusUseCase(repo domain.Repository) *ProcessOrderStatusUseCase {
	return &ProcessOrderStatusUseCase{repo: repo}
}

func (u *ProcessOrderStatusUseCase) Execute(ctx context.Context, input in.ProcessOrderStatusInput) error {
	ord, err := u.repo.GetOrderByID(ctx, input.OrderID)
	if err != nil {
		return &app.AppError{
			Code:    app.CodeNotFound,
			Message: domain.ErrOrderNotFound.Error(),
			Err:     err,
		}
	}

	if input.Amount != ord.Total {
		return &app.AppError{
			Code:    app.CodeConflict,
			Message: domain.ErrOrderAmountMismatch.Error(),
			Err:     domain.ErrOrderAmountMismatch,
		}
	}

	if err := ord.ChangeStatus(domain.OrderStatus(input.Status), time.Now().UTC()); err != nil {
		return &app.AppError{
			Code:    app.CodeConflict,
			Message: err.Error(),
			Err:     err,
		}
	}

	if err := u.repo.Update(ctx, ord); err != nil {
		return err
	}
	return nil
}
