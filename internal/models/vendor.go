package models

import "gorm.io/gorm"

type Vendor struct {
	gorm.Model
	Name   string `json:"name" gorm:"unique"`
	UserID uint   `json:"user_id"`
	User   User   `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
