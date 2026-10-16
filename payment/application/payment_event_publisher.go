package application

import (
	"context"

	"ops/payment/application/out"
)

type PaymentEventPublisher interface {
	PublishPaymentStatusChanged(ctx context.Context, event out.PaymentStatusChangedEvent) error
	PublishPaymentCreated(ctx context.Context, event out.PaymentCreatedEvent) error
}
