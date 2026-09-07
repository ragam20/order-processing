package event

type OrderCancelledEvent struct{}

func (e *OrderCancelledEvent) Handle() {}
