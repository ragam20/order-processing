package message

type OrderCanceled struct {
	OrderID int64 `json:"orderId"`
}

type OrderCompleted struct {
	OrderID int64 `json:"orderId"`
}
