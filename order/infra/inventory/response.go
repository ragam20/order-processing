package inventory

type CheckAvailabilityItemResponse struct {
	ProductID   int64  `json:"productId"`
	Name        string `json:"name"`
	Price       int64  `json:"price"`
	Quantity    int32  `json:"quantity"`
	Available   int    `json:"available"`
	IsAvailable bool   `json:"isAvailable"`
}

type CheckAvailabilityResponse struct {
	Success bool                            `json:"success"`
	Data    []CheckAvailabilityItemResponse `json:"data"`
}
