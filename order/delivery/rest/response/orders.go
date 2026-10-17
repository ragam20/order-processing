package response

import "time"

type OrderResponse struct {
	ID          int64               `json:"id"`
	Total       int64               `json:"total"`
	Status      string              `json:"status"`
	CancelledAt *time.Time          `json:"cancelledAt"`
	CompletedAt *time.Time          `json:"completedAt"`
	CreatedAt   time.Time           `json:"createdAt"`
	UpdatedAt   time.Time           `json:"updatedAt"`
	Items       []OrderItemResponse `json:"items"`
}

type OrderItemResponse struct {
	ID        int64     `json:"id"`
	ProductID int64     `json:"productId"`
	Name      string    `json:"name"`
	Price     int64     `json:"price"`
	Quantity  int32     `json:"quantity"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
