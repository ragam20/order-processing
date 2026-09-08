package inventory

type CheckAvailabilityRequest struct {
	Items []ItemRequest `json:"items"`
}

type ItemRequest struct {
	ProductID int64 `json:"productId"`
	Quantity  int32 `json:"quantity"`
}
