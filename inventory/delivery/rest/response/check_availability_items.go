package response

type ItemAvailabilityResponse struct {
	ProductID   int64  `json:"productId"`
	Name        string `json:"name"`
	Price       int64  `json:"price"`
	Quantity    int32  `json:"quantity"`
	Available   int    `json:"available"`
	IsAvailable bool   `json:"isAvailable"`
}
