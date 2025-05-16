package db

import (
	"github.com/Xebec19/lms/provision/internal/db/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.User{},
		&models.Role{},
		&models.Permission{},
	)

	if err != nil {
		return err
	}

	// Create default roles and permissions
	roles := []models.Role{
		{
			Name:        "admin",
			Description: "Administrator role with all permissions",
			Permissions: []models.Permission{
				{Name: "create:course", Description: "Create course"},
				{Name: "delete:course", Description: "Delete course"},
				{Name: "update:course", Description: "Update course"},
				{Name: "view:course", Description: "View course"},
			},
		},
		{
			Name:        "user",
			Description: "Regular user role with limited permissions",
			Permissions: []models.Permission{
				{Name: "view:course", Description: "View course"},
			},
		},
	}

	// Check if roles already exist
	for _, role := range roles {
		var count int64
		err := db.Model(&models.Role{}).Where("name = ?", role.Name).Count(&count).Error
		if err != nil {
			return err
		}
		if count == 0 {
			err := db.Create(&role).Error
			if err != nil {
				return err
			}
		}
	}

	return nil
}
