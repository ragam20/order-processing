package messaging

const (
	PaymentEventsExchange    = "payment.events"
	PaymentPaidRoutingKey    = "payment.paid"
	PaymentFailedRoutingKey  = "payment.failed"
	PaymentExpiredRoutingKey = "payment.expired"
	PaymentCreatedRoutingKey = "payment.created"
)
