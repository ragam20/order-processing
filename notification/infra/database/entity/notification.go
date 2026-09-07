package entity

import "time"

type NotificationEntity struct {
	ID      int64 `gorm:"primaryKey;autoIncrement"`
	Title   string
	Body    string
	Status  bool
	Channel string

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (NotificationEntity) TableName() string {
	return "notifications"
}
