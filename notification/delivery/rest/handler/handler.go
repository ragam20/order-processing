package handler

import "ops/notification/application/usecase"

type Handler interface{}

type handler struct {
	getNotificationsUseCase   *usecase.GetNotificationsUseCase
	createNotificationUseCase *usecase.CreateNotificationUseCase
}

func NewHandler(getNotificationsUseCase *usecase.GetNotificationsUseCase, createUseCase *usecase.CreateNotificationUseCase) Handler {
	return &handler{
		getNotificationsUseCase:   getNotificationsUseCase,
		createNotificationUseCase: createUseCase,
	}
}
