package usecase

import (
	"context"
	"time"

	"ops/notification/application/in"
	"ops/notification/domain"
	"ops/shared/app"
)

type CreateNotificationUseCase struct {
	repo domain.Repository
}

func NewCreateUseCase(repo domain.Repository) *CreateNotificationUseCase {
	return &CreateNotificationUseCase{repo: repo}
}

func (u *CreateNotificationUseCase) Execute(ctx context.Context, input in.CreateNotificationInput) *app.AppError {
	notification := &domain.Notification{
		ID:      time.Now().UnixMilli(),
		Title:   input.Title,
		Body:    input.Body,
		Channel: input.Channel,
		Status:  false,
	}
	if err := u.repo.Create(ctx, notification); err != nil {
		return &app.AppError{
			Code:    app.CodeInvalidInput,
			Message: domain.ErrCreateNotification.Error(),
			Err:     err,
		}
	}
	return nil
}
