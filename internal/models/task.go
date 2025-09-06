package models

import "time"

type Task struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    ProjectID uint      `gorm:"not null" json:"project_id"`
    Title     string    `gorm:"not null" json:"title"`
    Status    string    `gorm:"default:TODO" json:"status"` // TODO / IN_PROGRESS / DONE
    UserId uint     `json:"user_id"`
    CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}