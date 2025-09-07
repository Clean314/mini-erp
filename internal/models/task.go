package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title     string
	Status    string `gorm:"default:TODO"` // TODO, INPROGRESS, DONE
	ProjectID uint
	Assignee  uint
}