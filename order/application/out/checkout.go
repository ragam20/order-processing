package out

type CheckoutResult struct {
	OrderID int64
	Status  string
	Total   int64
	Items   []CheckoutItemResult
	Payment PaymentResult
}

type CheckoutItemResult struct {
	ProductID int64
	Name      string
	Price     int64
	Quantity  int32
}
