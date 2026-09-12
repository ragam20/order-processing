package usecase

import (
	"context"

	"ops/payment/application"
	"ops/payment/application/out"
	"ops/payment/domain"
	"ops/shared/app"
)

type GetPaymentStatus struct {
	provider application.PaymentProvider
}

func NewGetPaymentStatus(provider application.PaymentProvider) *GetPaymentStatus {
	return &GetPaymentStatus{provider: provider}
}

func (u *GetPaymentStatus) Execute(ctx context.Context, providerReference string) (*out.GetStatusPaymentResult, *app.AppError) {
	providerRes, err := u.provider.GetStatus(ctx, providerReference)
	if err != nil {
		return nil, &app.AppError{
			Code:    app.CodePaymentCreationFailed,
			Message: domain.ErrUnavailablePayment.Error(),
			Err:     err,
		}
	}

	return &out.GetStatusPaymentResult{
		ProviderReference:     providerRes.ProviderReference,
		Provider:              u.provider.GetProviderName(),
		ProviderTransactionID: providerRes.ProviderTransactionID,
		Status:                string(providerRes.Status),
		Amount:                providerRes.GrossAmount,
		PaidAt:                providerRes.PaidAt,
	}, nil
}
