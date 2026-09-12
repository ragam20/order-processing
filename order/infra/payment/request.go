package payment

type CreatePaymentRequest struct {
	OrderID     int64                `json:"orderId"`
	TotalAmount int64                `json:"totalAmount"`
	Items       []PaymentItemRequest `json:"items"`
	Customer    CustomerRequest      `json:"customer"`
	FinishURL   string               `json:"finishUrl"`
}

type CustomerRequest struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type PaymentItemRequest struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Price    int64  `json:"price"`
	Quantity int32  `json:"quantity"`
}
