package request

type CheckoutRequest struct {
	Items    []CheckoutRequestItem `json:"items"`
	Customer CustomerRequest       `json:"customer" binding:"required"`
}

type CheckoutRequestItem struct {
	ProductID int64 `json:"productId"`
	Quantity  int32 `json:"quantity"`
}

type CustomerRequest struct {
	ID    string `json:"id"`
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Phone string `json:"phone"`
}
