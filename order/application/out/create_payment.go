package out

type PaymentResult struct {
	Status                string
	ProviderReference     string
	ProviderTransactionID string
	CheckoutURL           string
	PaymentToken          *string
}
