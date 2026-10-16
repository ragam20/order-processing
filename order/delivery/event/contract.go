package event

const (
	PaymentEventsExchange = "payment.events"
	QueueName             = "order.payment-status"

	PaymentPaidRoutingKey    = "payment.paid"
	PaymentFailedRoutingKey  = "payment.failed"
	PaymentExpiredRoutingKey = "payment.expired"
)
