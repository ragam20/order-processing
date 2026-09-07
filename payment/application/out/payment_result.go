package out

import "encoding/json"

type PaymentResult struct {
	ProviderReference     string
	ProviderTransactionID string
	CheckoutURL           string
	PaymentToken          string
	RawResponse           json.RawMessage
}
