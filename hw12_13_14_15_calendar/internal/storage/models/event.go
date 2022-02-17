package storagemodels

import "time"

type Event struct {
	ID          int           `json:"id"`
	Title       string        `json:"title"`
	Duration    time.Duration `json:"duration"`
	Description string        `json:"description"`
	NotifyAt    time.Time     `json:"notifyAt"`
	CreatedAt   time.Time     `json:"createdAt"`
	UpdatedAt   time.Time     `json:"updatedAt"`
}
