package db

import (
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	// err := db.AutoMigrate(
	// 	&models.User{},
	// 	&models.Role{},
	// 	&models.Permission{},
	// )

	// if err != nil {
	// 	return err
	// }

	return nil
}
