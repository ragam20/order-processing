package usecase

import (
	"context"
	"errors"

	"ops/notification/domain"
	"ops/shared/app"
)

type ReadNotificationUseCase struct {
	repo domain.Repository
}

func NewReadNotificationUseCase(repo domain.Repository) *ReadNotificationUseCase {
	return &ReadNotificationUseCase{repo: repo}
}

func (u *ReadNotificationUseCase) Execute(ctx context.Context, notificationID int64) *app.AppError {
	n, err := u.repo.GetByID(ctx, notificationID)
	if err != nil {
		if errors.Is(err, domain.ErrNoficationNotFound) {
			return &app.AppError{
				Code:    app.CodeNotFound,
				Message: domain.ErrNoficationNotFound.Error(),
				Err:     err,
			}
		}

		return &app.AppError{
			Code:    app.CodeInternalError,
			Message: "internal error",
			Err:     err,
		}
	}

	n.Read()

	if err := u.repo.Update(ctx, n); err != nil {
		if errors.Is(err, domain.ErrUpdateNotification) {
			return &app.AppError{
				Code:    app.CodeConflict,
				Message: domain.ErrUpdateNotification.Error(),
				Err:     err,
			}
		}
		return &app.AppError{
			Code:    app.CodeInternalError,
			Message: "internal error",
			Err:     err,
		}
	}

	return nil
}
