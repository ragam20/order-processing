package usecase

import (
	"context"
	"errors"

	"ops/order/domain"
	"ops/shared/app"
)

type CancelOrderUseCase struct {
	repo domain.Repository
}

func NewCancelOrderUseCase(repo domain.Repository) *CancelOrderUseCase {
	return &CancelOrderUseCase{
		repo: repo,
	}
}

func (u *CancelOrderUseCase) Execute(ctx context.Context, orderID int64) *app.AppError {
	order, err := u.repo.GetOrderByID(ctx, orderID)
	if err != nil {
		if errors.Is(err, domain.ErrOrderNotFound) {
			return &app.AppError{
				Code:    app.CodeNotFound,
				Message: domain.ErrOrderNotFound.Error(),
				Err:     err,
			}
		}
		return &app.AppError{
			Code:    app.CodeInternalError,
			Message: "internal error",
			Err:     err,
		}
	}

	if err := order.Cancel(); err != nil {
		return &app.AppError{
			Code:    app.CodeConflict,
			Message: domain.ErrOrderCanceled.Error(),
			Err:     err,
		}
	}

	if err := u.repo.Update(ctx, order); err != nil {
		return &app.AppError{
			Code:    app.CodeConflict,
			Message: domain.ErrOrderUpdate.Error(),
			Err:     err,
		}
	}

	return nil
}
