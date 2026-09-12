package application

import (
	"context"

	"ops/payment/application/in"
	"ops/payment/application/out"
)

type PaymentProvider interface {
	GetProviderName() string
	CreatePayment(ctx context.Context, input in.CreatePaymentInput) (out.PaymentResult, error)
	GetStatus(ctx context.Context, reference string) (out.PaymentStatusResult, error)
	ParseWebhook(input in.ParseWebhook) (out.PaymentNotificationResult, error)
}
