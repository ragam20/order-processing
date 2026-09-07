package usecase

import (
	"context"
	"errors"

	"ops/order/domain"
	"ops/shared/app"
)

type CompletedOrderUseCase struct {
	repo domain.Repository
}

func NewSuccessOrderUseCase(repo domain.Repository) *CompletedOrderUseCase {
	return &CompletedOrderUseCase{repo: repo}
}

func (u *CompletedOrderUseCase) Execute(ctx context.Context, orderID int64) error {
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

	if err := order.Completed(); err != nil {
		return &app.AppError{
			Code:    app.CodeConflict,
			Message: domain.ErrOrderSuccess.Error(),
			Err:     err,
		}
	}

	return nil
}
