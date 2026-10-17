package usecase

import (
	"context"

	"ops/order/application"
	"ops/order/application/out"
	"ops/order/domain"
	"ops/shared/app"
)

type GetOrdersByStatusUseCase struct {
	repo application.OrderQueryRepository
}

func NewGetOrdersByStatusUseCase(repo application.OrderQueryRepository) *GetOrdersByStatusUseCase {
	return &GetOrdersByStatusUseCase{repo: repo}
}

func (u *GetOrdersByStatusUseCase) Execute(ctx context.Context, status domain.OrderStatus) ([]out.OrderWithItemsResult, *app.AppError) {
	if !isValidOrderStatus(status) {
		return nil, &app.AppError{
			Code:    app.CodeInvalidInput,
			Message: "invalid order status",
		}
	}

	orders, err := u.repo.GetOrdersByStatus(ctx, status)
	if err != nil {
		return nil, &app.AppError{
			Code:    app.CodeInternalError,
			Message: "failed to get orders by status",
			Err:     err,
		}
	}

	return orders, nil
}

func isValidOrderStatus(status domain.OrderStatus) bool {
	switch status {
	case domain.OrderPending, domain.OrderProcess, domain.OrderCancelled, domain.OrderSuccess:
		return true
	default:
		return false
	}
}
