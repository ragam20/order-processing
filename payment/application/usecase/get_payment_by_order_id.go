package usecase

import (
	"context"

	"ops/payment/application/out"
	"ops/payment/domain"
	"ops/shared/app"
)

type GetPaymentByOrderIDUseCase struct {
	repo domain.Repository
}

func NewGetPaymentByOrderIDUseCase(repo domain.Repository) *GetPaymentByOrderIDUseCase {
	return &GetPaymentByOrderIDUseCase{repo: repo}
}

func (u *GetPaymentByOrderIDUseCase) Execute(ctx context.Context, orderID int64) (*out.PaymentDetailResult, *app.AppError) {
	if orderID <= 0 {
		return nil, &app.AppError{
			Code:    app.CodeInvalidRequest,
			Message: "order id must be greater than zero",
		}
	}

	payment, err := u.repo.GetByOrderID(ctx, orderID)
	if err != nil {
		if err == domain.ErrPaymentNotFound {
			return nil, &app.AppError{
				Code:    app.CodeNotFound,
				Message: domain.ErrPaymentNotFound.Error(),
				Err:     err,
			}
		}

		return nil, &app.AppError{
			Code:    app.CodeInternalError,
			Message: "failed to get payment",
			Err:     err,
		}
	}

	return &out.PaymentDetailResult{
		ID:                    payment.ID,
		OrderID:               payment.OrderID,
		Amount:                payment.Amount,
		Status:                string(payment.Status),
		Provider:              payment.Provider,
		ProviderReference:     payment.ProviderReference,
		ProviderTransactionID: payment.ProviderTransactionID,
		CheckoutURL:           payment.CheckoutURL,
		PaymentToken:          payment.PaymentToken,
		PaidAt:                payment.PaidAt,
		CreatedAt:             payment.CreatedAt,
		UpdatedAt:             payment.UpdatedAt,
	}, nil
}
