package domain

import "errors"

var (
	ErrOrderNotFound     = errors.New("order not found")
	ErrOrderCanceled     = errors.New("order can't be canceled")
	ErrOrderUpdate       = errors.New("order can't be updated")
	ErrOrderSuccess      = errors.New("order can't be successed")
	ErrInsufficientItems = errors.New("insufficient items")
)
