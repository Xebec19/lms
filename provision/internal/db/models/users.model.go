package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string `json:"first_name" gorm:"type:varchar(100);not null"`
	LastName  string `json:"last_name" gorm:"type:varchar(100)"`
	Email     string `json:"email" gorm:"type:varchar(100);unique;not null"`
	Password  string `json:"password" gorm:"type:varchar(100);not null"`
	Role      string `json:"role" gorm:"type:varchar(50);not null;default:'user'"`
	Phone     string `json:"phone" gorm:"type:varchar(15)"`
}
