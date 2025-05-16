package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name        string       `json:"name" gorm:"type:varchar(100);not null"`
	Description string       `json:"description" gorm:"type:text"`
	Permissions []Permission `json:"permissions" gorm:"many2many:role_permissions;"`
}
