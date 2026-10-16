package in

type Customer struct {
	ID    string
	Name  string
	Email string
	Phone string
}

type Item struct {
	ID       string
	Name     string
	Price    int64
	Quantity int32
}

type CreatePaymentInput struct {
	OrderID   int64
	Amount    int64
	Items     []Item
	Customer  Customer
	FinishURL string
}
