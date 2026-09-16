package event

import (
	"context"

	"ops/order/application/usecase"
	"ops/order/delivery/event/message"
	"ops/shared/rmq"
)

type ProcessOrderStatusChangedEvent struct {
	uc       *usecase.ProcessOrderStatusUseCase
	consumer rmq.Consumer
}

func NewProcessOrderStatusChangedEvent(uc *usecase.ProcessOrderStatusUseCase, consumer rmq.Consumer) *ProcessOrderStatusChangedEvent {
	return &ProcessOrderStatusChangedEvent{uc: uc, consumer: consumer}
}

func (e *ProcessOrderStatusChangedEvent) Handle(ctx context.Context, msg message.PaymentStatusChanged) {
}
