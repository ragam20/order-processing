package messaging

import (
	"context"
	"encoding/json"
	"fmt"

	"ops/payment/application"
	"ops/payment/application/out"
	"ops/payment/domain"
	"ops/shared/event"
	"ops/shared/rmq"
)

type paymentEventPublisher struct {
	publisher rmq.Publisher
}

func NewPaymentEventPublisher(publisher rmq.Publisher) application.PaymentEventPublisher {
	return &paymentEventPublisher{publisher: publisher}
}

func (p *paymentEventPublisher) PublishPaymentStatusChanged(ctx context.Context, paymentEvent out.PaymentStatusChangedEvent) error {
	routingKey, err := routingKeyFor(domain.PaymentStatus(paymentEvent.Status))
	if err != nil {
		return err
	}

	data, err := json.Marshal(paymentEvent)
	if err != nil {
		return err
	}

	return p.publisher.Publish(
		ctx,
		PaymentEventsExchange,
		event.ExchangeTypeTopic,
		routingKey,
		data,
	)
}

func routingKeyFor(status domain.PaymentStatus) (string, error) {
	switch status {
	case domain.PaymentStatusPaid:
		return PaymentPaidRoutingKey, nil
	case domain.PaymentStatusFailed:
		return PaymentFailedRoutingKey, nil
	case domain.PaymentStatusExpired:
		return PaymentExpiredRoutingKey, nil
	default:
		return "", fmt.Errorf("unsupported payment event status: %s", status)
	}
}
