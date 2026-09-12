package usecase

import (
	"context"

	"ops/payment/application"
	"ops/payment/application/in"
	"ops/payment/application/out"
	"ops/payment/domain"
	"ops/shared/app"
)

type CreatePaymentUseCase struct {
	repo     domain.Repository
	provider application.PaymentProvider
}

func NewCreatePaymentUseCase(repo domain.Repository, provider application.PaymentProvider) *CreatePaymentUseCase {
	return &CreatePaymentUseCase{repo: repo, provider: provider}
}

func (u *CreatePaymentUseCase) Execute(ctx context.Context, input in.CreatePaymentInput) (out.PaymentResult, *app.AppError) {
	providerRes, err := u.provider.CreatePayment(ctx, input)
	if err != nil {
		return out.PaymentResult{}, &app.AppError{
			Code:    app.CodePaymentCreationFailed,
			Message: domain.ErrUnavailablePayment.Error(),
			Err:     err,
		}
	}

	if err := u.repo.Create(ctx, &domain.Payment{
		Status:                domain.PaymentStatusPending,
		Provider:              u.provider.GetProviderName(),
		ProviderReference:     providerRes.ProviderReference,
		ProviderTransactionID: providerRes.ProviderTransactionID,
		CheckoutURL:           providerRes.CheckoutURL,
		PaymentToken:          providerRes.PaymentToken,
		ResponseProvider:      providerRes.RawResponse,
	}, input.OrderID); err != nil {
		return out.PaymentResult{}, &app.AppError{
			Code:    app.CodeConflict,
			Message: domain.ErrCreatePayment.Error(),
			Err:     err,
		}
	}

	providerRes.Status = string(domain.PaymentStatusPending)

	return providerRes, nil
}
