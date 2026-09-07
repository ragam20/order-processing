package usecase

import (
	"context"

	"ops/notification/domain"
)

type GetNotificationsUseCase struct {
	repo domain.Repository
}

func NewGetNotificationsUseCase(repo domain.Repository) *GetNotificationsUseCase {
	return &GetNotificationsUseCase{repo: repo}
}

func (u *GetNotificationsUseCase) Execute(ctx context.Context) []domain.Notification {
	return u.repo.GetNotifications(ctx)
}
