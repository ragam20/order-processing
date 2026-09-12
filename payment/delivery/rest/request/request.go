package request

type CreatePaymentRequest struct {
	OrderID     int64                `json:"orderId" binding:"required"`
	TotalAmount int64                `json:"totalAmount" binding:"required"`
	Items       []PaymentItemRequest `json:"items" binding:"required,min=1"`
	Customer    CustomerRequest      `json:"customer" binding:"required"`
	FinishURL   string               `json:"finishUrl"`
}

type CustomerRequest struct {
	ID    string `json:"id"`
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Phone string `json:"phone"`
}

type PaymentItemRequest struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Price    int64  `json:"price"`
	Quantity int32  `json:"quantity"`
}
