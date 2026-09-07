package handler

import "ops/payment/application/usecase"

type Handler interface{}

type handler struct {
	processPaymentUseCase *usecase.ProcessPaymentUseCase
}

func NewHandler(processPaymentUseCase *usecase.ProcessPaymentUseCase) Handler {
	return &handler{processPaymentUseCase: processPaymentUseCase}
}
