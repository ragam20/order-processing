package request

type CheckAvailabilityRequest struct {
	Items []CheckAvailabilityItem `json:"items"`
}

type CheckAvailabilityItem struct {
	ProductID int64 `json:"productId"`
	Quantity  int32 `json:"quantity"`
}
