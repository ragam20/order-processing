package message

type OrderCanceled struct {
	OrderID int64 `json:"orderId"`
}

type OrderCompleted struct {
	OrderID int64 `json:"orderId"`
}

type PaymentStatusChanged struct {
	OrderID int64  `json:"orderId"`
	Status  string `json:"status"`
	Amount  int64  `json:"amount"`
}
