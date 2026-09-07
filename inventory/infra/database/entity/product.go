package entity

import "time"

type ProductEntity struct {
	ID        int64 `gorm:"autoIncrement;primaryKey"`
	Name      string
	ImageURL  string
	Available int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (ProductEntity) TableName() string {
	return "products"
}
