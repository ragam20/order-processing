package domain

import "errors"

var (
	ErrCreatePayment        = errors.New("failed to create payment")
	ErrUnavailablePayment   = errors.New("payment provider is temporarily unavailable")
	ErrPaymentNotFound      = errors.New("payment not found")
	ErrInvalidPaymentAmount = errors.New("invalid payment amount")
)
