package event

import "time"

// Creating an event entity belonging to a specific user
type Event struct {
	ID      int64     `gorm:"primaryKey" json:"id"`
	OwnerId int64     `gorm:"not null" json:"owner_id"`
	Name    string    `json:"name" binding:"required,min=1"`
	Content string    `json:"content"`
	Time    time.Time `json:"time" binding:"required"`
}
