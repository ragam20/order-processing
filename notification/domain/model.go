package domain

import "time"

type Notification struct {
	ID      int64
	Title   string
	Body    string
	Status  bool
	Channel string

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (n *Notification) Read() {
	n.Status = true
}
