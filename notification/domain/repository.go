package domain

import "context"

type Repository interface {
	Create(ctx context.Context, notification *Notification) error
	Update(ctx context.Context, notification *Notification) error
	GetByID(ctx context.Context, notificationID int64) (*Notification, error)
	GetNotifications(ctx context.Context) []Notification
}
