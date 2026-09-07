package in

type CheckoutItemInput struct {
	ItemID   int64
	Quantity int32
}

type CheckoutInput struct {
	Items []CheckoutItemInput
}
