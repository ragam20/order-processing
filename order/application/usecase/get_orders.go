package usecase

import (
	"context"

	"ops/order/application"
	"ops/order/application/out"
	"ops/shared/app"
)

type GetOrdersUseCase struct {
	repo application.OrderQueryRepository
}

func NewGetOrdersUseCase(repo application.OrderQueryRepository) *GetOrdersUseCase {
	return &GetOrdersUseCase{repo: repo}
}

func (u *GetOrdersUseCase) Execute(ctx context.Context) ([]out.OrderWithItemsResult, *app.AppError) {
	orders, err := u.repo.GetOrdersWithItems(ctx)
	if err != nil {
		return nil, &app.AppError{
			Code:    app.CodeInternalError,
			Message: "failed to get orders",
			Err:     err,
		}
	}

	return orders, nil
}
