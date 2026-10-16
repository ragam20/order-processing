package request

type CreateProductsRequest struct {
	Items []CreateProductItemRequest `json:"items"`
}

type CreateProductItemRequest struct {
	Name      string `json:"name"`
	ImageURL  string `json:"imageUrl"`
	Price     int64  `json:"price"`
	Available int    `json:"available"`
}
