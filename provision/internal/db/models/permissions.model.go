package models

import "gorm.io/gorm"

type Permission struct {
	gorm.Model
	Name        string `json:"name" gorm:"type:varchar(100);not null"`
	Description string `json:"description" gorm:"type:text"`
}
