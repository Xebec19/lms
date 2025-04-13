package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	UserId      uuid.UUID `gorm:"type:uuid;uuid_generate:uuid_generate_v4();primaryKey"`
	Username    string    `gorm:"not null;unique"`
	Email       string    `gorm:"not null;unique"`
	Password    string    `gorm:"not null"`
	Role        string    `gorm:"not null"`
	UserPicture string    `gorm:"not null"`
}
