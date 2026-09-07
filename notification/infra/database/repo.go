package database

import (
	"context"
	"errors"
	"ops/notification/domain"
	"ops/notification/infra/database/entity"
	"time"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) domain.Repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, notification *domain.Notification) error {
	if err := r.db.WithContext(ctx).Create(&notification).Error; err != nil {
		return err
	}
	return nil
}

func (r *repository) Update(ctx context.Context, notification *domain.Notification) error {
	e := &entity.NotificationEntity{
		ID:        notification.ID,
		Title:     notification.Title,
		Body:      notification.Body,
		Status:    notification.Status,
		Channel:   notification.Channel,
		CreatedAt: notification.CreatedAt,
		UpdatedAt: time.Now(),
	}
	if err := r.db.WithContext(ctx).Model(&entity.NotificationEntity{}).Updates(&e).Error; err != nil {
		return domain.ErrUpdateNotification
	}

	return nil
}

func (r *repository) GetByID(ctx context.Context, notificationID int64) (*domain.Notification, error) {
	var e entity.NotificationEntity
	if err := r.db.WithContext(ctx).Where("id = ?", notificationID).First(&e).Error; err != nil {
		if errors.Is(err, domain.ErrNoficationNotFound) {
			return nil, domain.ErrNoficationNotFound
		}

		return nil, err
	}

	d := &domain.Notification{
		ID:        e.ID,
		Title:     e.Title,
		Body:      e.Body,
		Status:    e.Status,
		Channel:   e.Channel,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}

	return d, nil
}

func (r *repository) GetNotifications(ctx context.Context) []domain.Notification {
	var e []entity.NotificationEntity
	if err := r.db.WithContext(ctx).Find(&e).Error; err != nil {
		return []domain.Notification{}
	}

	d := make([]domain.Notification, 0)
	for _, en := range e {
		d = append(d, domain.Notification{
			ID:      en.ID,
			Title:   en.Title,
			Body:    en.Body,
			Status:  en.Status,
			Channel: en.Channel,

			CreatedAt: en.CreatedAt,
			UpdatedAt: en.UpdatedAt,
		})
	}

	return d
}
