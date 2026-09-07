package domain

import "errors"

var (
	ErrCreateNotification = errors.New("error create notification")
	ErrUpdateNotification = errors.New("can't update notification")
	ErrNoficationNotFound = errors.New("notification not found")
)
