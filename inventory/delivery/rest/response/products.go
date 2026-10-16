package response

import "time"

type ProductResponse struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	ImageURL  string    `json:"imageUrl"`
	Price     int64     `json:"price"`
	Available int       `json:"available"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
