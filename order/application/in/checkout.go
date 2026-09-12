package in

type CheckoutItemInput struct {
	ProductID int64
	Quantity  int32
}

type CheckoutInput struct {
	Items     []CheckoutItemInput
	Customer  CustomerInput
	FinishURL string
}

type CustomerInput struct {
	ID    string
	Name  string
	Email string
	Phone string
}
