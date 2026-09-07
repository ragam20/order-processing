package usecase

import (
	"context"
	"errors"

	"ops/order/application/out"
	"ops/order/domain"
	"ops/shared/app"
)

type GetDetailOrderUseCase struct {
	repo domain.Repository
}

func NewGetDetailOrderUseCase(repo domain.Repository) *GetDetailOrderUseCase {
	return &GetDetailOrderUseCase{
		repo: repo,
	}
}

func (u *GetDetailOrderUseCase) Execute(ctx context.Context, orderID int64) (*out.OrderDetailResult, error) {
	order, err := u.repo.GetOrderByID(ctx, orderID)
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

	orderItem := u.repo.GetItems(ctx, orderID)

	return &out.OrderDetailResult{
		Order: order,
		Items: orderItem,
	}, nil
}
