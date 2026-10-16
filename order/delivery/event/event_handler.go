package event

import (
	"context"
	"encoding/json"
	"fmt"

	"ops/order/application/in"
	"ops/order/application/usecase"
	"ops/order/delivery/event/message"
	"ops/order/domain"

	amqp "github.com/rabbitmq/amqp091-go"
)

type EventHandler struct {
	processOrderStatusUC *usecase.ProcessOrderStatusUseCase
}

func NewEventHandler(processOrderStatusUC *usecase.ProcessOrderStatusUseCase) *EventHandler {
	return &EventHandler{processOrderStatusUC}
}

func (h *EventHandler) ProcessOrderStatus(ctx context.Context, delivery amqp.Delivery) error {
	var msg message.PaymentStatusChanged

	if err := json.Unmarshal(delivery.Body, &msg); err != nil {
		return err
	}

	if msg.OrderID <= 0 {
		return fmt.Errorf("invalid order id: %d", msg.OrderID)
	}
	if msg.Amount < 0 {
		return fmt.Errorf("invalid payment amount: %d", msg.Amount)
	}

	var status domain.OrderStatus
	switch delivery.RoutingKey {
	case PaymentPaidRoutingKey:
		status = domain.OrderSuccess
	case PaymentFailedRoutingKey:
		status = domain.OrderCancelled
	case PaymentExpiredRoutingKey:
		status = domain.OrderCancelled
	default:
		return fmt.Errorf("unsupported routing key: %s", delivery.RoutingKey)
	}

	return h.processOrderStatusUC.Execute(ctx, in.ProcessOrderStatusInput{
		OrderID: msg.OrderID,
		Status:  string(status),
		Amount:  msg.Amount,
	})
}
