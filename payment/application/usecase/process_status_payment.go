package usecase

import (
	"context"

	"ops/payment/application"
	"ops/payment/application/in"
	"ops/payment/application/out"
	"ops/payment/domain"
)

type ProcessStatusPaymentUseCase struct {
	repo      domain.Repository
	publisher application.PaymentEventPublisher
}

func NewProcessStatusPaymentUseCase(
	repo domain.Repository,
	publisher application.PaymentEventPublisher,
) *ProcessStatusPaymentUseCase {
	return &ProcessStatusPaymentUseCase{repo: repo, publisher: publisher}
}

func (u *ProcessStatusPaymentUseCase) Execute(ctx context.Context, input in.ProcessPaymentInput) error {
	payment, err := u.repo.GetByProviderReference(ctx, input.ProviderReference)
	if err != nil {
		return err
	}

	if payment.Amount != input.GrossAmount {
		return domain.ErrInvalidPaymentAmount
	}

	if payment.Status == input.Status {
		return nil
	}

	payment.Status = input.Status
	payment.ProviderTransactionID = input.ProviderTransactionID
	payment.ResponseProvider = input.RawPayload
	payment.PaidAt = input.PaidAt

	if err := u.repo.Update(ctx, payment); err != nil {
		return err
	}

	switch input.Status {
	case domain.PaymentStatusPaid, domain.PaymentStatusFailed, domain.PaymentStatusExpired:
		return u.publisher.PublishPaymentStatusChanged(ctx, out.PaymentStatusChangedEvent{
			OrderID:               payment.OrderID,
			Status:                string(input.Status),
			Amount:                payment.Amount,
			Provider:              payment.Provider,
			ProviderReference:     payment.ProviderReference,
			ProviderTransactionID: payment.ProviderTransactionID,
		})
	default:
		return nil
	}
}
