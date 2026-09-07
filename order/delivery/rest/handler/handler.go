package handler

import "ops/order/application/usecase"

type Handler interface{}

type handler struct {
	cancelOrderUseCase    *usecase.CancelOrderUseCase
	getDetailOrderUseCase *usecase.GetDetailOrderUseCase
}

func NewHandler(cancelOrderUseCase *usecase.CancelOrderUseCase, getDetailOrderUseCase *usecase.GetDetailOrderUseCase) Handler {
	return &handler{
		cancelOrderUseCase:    cancelOrderUseCase,
		getDetailOrderUseCase: getDetailOrderUseCase,
	}
}
