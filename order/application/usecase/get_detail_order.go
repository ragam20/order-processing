package usecase

import (
	"context"
	"errors"

	"ops/order/application"
	"ops/order/application/out"
	"ops/order/domain"
	"ops/shared/app"
)

type GetDetailOrderUseCase struct {
	repo application.OrderQueryRepository
}

func NewGetDetailOrderUseCase(repo application.OrderQueryRepository) *GetDetailOrderUseCase {
	return &GetDetailOrderUseCase{
		repo: repo,
	}
}

func (u *GetDetailOrderUseCase) Execute(ctx context.Context, orderID int64) (*out.OrderWithItemsResult, error) {
	order, err := u.repo.GetOrderDetail(ctx, orderID)
	if err != nil {
		if errors.Is(err, domain.ErrOrderNotFound) {
			return nil, &app.AppError{
				Code:    app.CodeNotFound,
				Message: domain.ErrOrderNotFound.Error(),
				Err:     err,
			}
		}
		return nil, &app.AppError{
			Code:    app.CodeInternalError,
			Message: "internal error",
			Err:     err,
		}
	}

	return order, nil
}
