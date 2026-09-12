package response

type CheckoutResponse struct {
	OrderID int64                  `json:"orderId"`
	Status  string                 `json:"status"`
	Total   int64                  `json:"total"`
	Items   []CheckoutItemResponse `json:"items"`
	Payment OrderPaymentResponse   `json:"payment"`
}

type CheckoutItemResponse struct {
	ProductID int64  `json:"productId"`
	Name      string `json:"name"`
	Price     int64  `json:"price"`
	Quantity  int32  `json:"quantity"`
}

type OrderPaymentResponse struct {
	Status       string  `json:"status"`
	CheckoutURL  string  `json:"checkoutUrl"`
	PaymentToken *string `json:"paymentToken"`
}
