package domain

import (
	"time"
)

type Product struct {
	ID        int64
	Name      string
	ImageURL  string
	Available int
	Price     int64
	CreatedAt time.Time
	UpdatedAt time.Time
}
