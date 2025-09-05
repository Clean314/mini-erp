package models

import "time"

type Project struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"not null" json:"name"`
	Description string    `json:"description"`
	OwnerID     uint      `json:"owner_id"`
	CreatedAt   time.Time `json:"created_at"`
}
