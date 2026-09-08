package domain

import "errors"

var (
	ErrProductNotFound   = errors.New("product not found")
	ErrProductUpdate     = errors.New("failed update product")
	ErrInsufficientStock = errors.New("insufficient stock")
	ErrProductQuantity   = errors.New("quantity must be greater than 0")
)
