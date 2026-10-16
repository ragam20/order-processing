package domain

import "errors"

var (
	ErrProductNotFound   = errors.New("product not found")
	ErrProductUpdate     = errors.New("failed update product")
	ErrInsufficientStock = errors.New("insufficient stock")
	ErrProductQuantity   = errors.New("quantity must be greater than 0")
	ErrProductName       = errors.New("product name must not be empty")
	ErrProductPrice      = errors.New("product price must not be negative")
	ErrProductAvailable  = errors.New("product available must not be negative")
)
