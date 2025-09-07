package models

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Name   string
	Owner  uint
	Tasks  []Task
}