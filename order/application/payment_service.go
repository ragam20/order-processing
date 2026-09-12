package application

import (
	"context"

	"ops/order/application/in"
	"ops/order/application/out"
)

type PaymentService interface {
	CreatePayment(ctx context.Context, input in.CreatePaymentInput) (out.PaymentResult, error)
	GetStatusPayment(ctx context.Context, providerReference string) (out.GetStatusPaymentResult, error)
}
