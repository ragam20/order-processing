package domain

import "errors"

var (
	ErrOrderNotFound     = errors.New("order not found")
	ErrOrderCanceled     = errors.New("order can't be canceled")
	ErrOrderUpdate       = errors.New("order can't be updated")
	ErrOrderSuccess      = errors.New("order can't be successed")
	ErrInsufficientItems = errors.New("insufficient items")
	ErrItemsEmpty        = errors.New("order items must not be empty")
	ErrInvalidQuantity   = errors.New("quantity must be greater than 0")
	ErrInvalidPrice      = errors.New("price must not be negative")
	ErrTotalOverflow     = errors.New("order total is too large")
)
