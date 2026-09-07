package domain

import "time"

type Product struct {
	ID        int64
	Name      string
	ImageURL  string
	Available int
	CreatedAt time.Time
	UpdatedAt time.Time
}
