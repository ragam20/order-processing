package usecase

import (
	"context"
	"fmt"
	"math"

	"ops/order/application"
	"ops/order/application/in"
	"ops/order/application/out"
	"ops/order/domain"
	"ops/shared/app"
)

type CheckoutUseCase struct {
	repo         domain.Repository
	inventorySvc application.InventoryService
	paymentSvc   application.PaymentService
}

func NewCheckoutUseCase(repo domain.Repository, inventorySvc application.InventoryService, paymentSvc application.PaymentService) *CheckoutUseCase {
	return &CheckoutUseCase{
		repo:         repo,
		inventorySvc: inventorySvc,
		paymentSvc:   paymentSvc,
	}
}

func (u *CheckoutUseCase) Execute(ctx context.Context, input in.CheckoutInput) (*out.CheckoutResult, *app.AppError) {
	if len(input.Items) == 0 {
		return nil, &app.AppError{
			Code:    app.CodeInvalidInput,
			Message: domain.ErrItemsEmpty.Error(),
			Err:     domain.ErrItemsEmpty,
		}
	}

	seen := make(map[int64]struct{}, len(input.Items))
	for _, item := range input.Items {
		if item.Quantity <= 0 {
			return nil, &app.AppError{
				Code:    app.CodeInvalidInput,
				Message: domain.ErrInvalidQuantity.Error(),
				Err:     domain.ErrInvalidQuantity,
			}
		}
		if _, exists := seen[item.ProductID]; exists {
			return nil, &app.AppError{
				Code:    app.CodeInvalidInput,
				Message: "duplicate product id",
			}
		}
		seen[item.ProductID] = struct{}{}
	}

	items, err := u.inventorySvc.CheckAvailabilityItem(ctx, input)
	if err != nil {
		return nil, &app.AppError{
			Code:    app.CodeConflict,
			Message: domain.ErrInsufficientItems.Error(),
			Err:     err,
		}
	}

	orderItem := make([]*domain.OrderItem, 0, len(items))
	var total int64
	for _, v := range items {
		if !v.IsAvailable {
			return nil, &app.AppError{
				Code:    app.CodeConflict,
				Message: domain.ErrInsufficientItems.Error(),
			}
		}

		if v.Price < 0 {
			return nil, &app.AppError{
				Code:    app.CodeInvalidInput,
				Message: domain.ErrInvalidPrice.Error(),
				Err:     domain.ErrInvalidPrice,
			}
		}

		quantity := int64(v.Quantity)
		if v.Price != 0 && quantity > math.MaxInt64/v.Price {
			return nil, &app.AppError{
				Code:    app.CodeConflict,
				Message: domain.ErrTotalOverflow.Error(),
				Err:     domain.ErrTotalOverflow,
			}
		}

		lineTotal := v.Price * quantity
		if total > math.MaxInt64-lineTotal {
			return nil, &app.AppError{
				Code:    app.CodeConflict,
				Message: domain.ErrTotalOverflow.Error(),
				Err:     domain.ErrTotalOverflow,
			}
		}

		total += lineTotal
		orderItem = append(orderItem, &domain.OrderItem{
			ProductID: v.ProductID,
			Name:      v.Name,
			Price:     v.Price,
			Quantity:  v.Quantity,
		})
	}

	order := &domain.Order{
		Total:  total,
		Status: domain.OrderPending,
	}

	_, err = u.repo.CreateWithItems(ctx, order, orderItem)
	if err != nil {
		return nil, &app.AppError{
			Code:    app.CodeInternalError,
			Message: "internal error",
			Err:     err,
		}
	}

	paymentItems := make([]in.PaymentItemInput, 0, len(orderItem))
	for _, item := range orderItem {
		paymentItems = append(paymentItems, in.PaymentItemInput{
			ID:       fmt.Sprintf("%d", item.ProductID),
			Name:     item.Name,
			Price:    item.Price,
			Quantity: item.Quantity,
		})
	}

	paymentResult, paymentErr := u.paymentSvc.CreatePayment(ctx, in.CreatePaymentInput{
		OrderID:     order.ID,
		TotalAmount: order.Total,
		Items:       paymentItems,
		Customer:    input.Customer,
	})
	if paymentErr != nil {
		paymentResult = out.PaymentResult{Status: "INITIATION_FAILED"}
	}

	resItem := make([]out.CheckoutItemResult, 0)
	for _, v := range orderItem {
		resItem = append(resItem, out.CheckoutItemResult{
			ProductID: v.ProductID,
			Name:      v.Name,
			Price:     v.Price,
			Quantity:  v.Quantity,
		})
	}

	return &out.CheckoutResult{
		OrderID: order.ID,
		Status:  string(order.Status),
		Total:   order.Total,
		Items:   resItem,
		Payment: out.PaymentResult{
			Status:       paymentResult.Status,
			CheckoutURL:  paymentResult.CheckoutURL,
			PaymentToken: paymentResult.PaymentToken,
		},
	}, nil
}
