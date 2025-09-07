package models

import "gorm.io/gorm"

type Approval struct {
	gorm.Model
	ProcessType string
	Status      string `gorm:"default:REQUESTED"` // REQUESTED, APPROVED, REJECTED
	RequestedBy uint
	ApprovedBy  *uint
}
