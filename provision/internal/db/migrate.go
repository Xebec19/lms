package db

import (
	"github.com/Xebec19/lms/provision/internal/db/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.User{},
	)
}
