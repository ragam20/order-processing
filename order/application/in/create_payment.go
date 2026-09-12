package in

type CreatePaymentInput struct {
	OrderID     int64
	TotalAmount int64
	Items       []PaymentItemInput
	Customer    CustomerInput
	FinishURL   string
}

type PaymentItemInput struct {
	ID       string
	Name     string
	Price    int64
	Quantity int32
}
