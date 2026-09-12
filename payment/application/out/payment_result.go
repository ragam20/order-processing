package out

import "encoding/json"

type PaymentResult struct {
	Status                string
	ProviderReference     string
	ProviderTransactionID string
	CheckoutURL           string
	PaymentToken          string
	RawResponse           json.RawMessage
}
