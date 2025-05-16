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

	// Step 1: Predefine all permissions
	allPermissions := []models.Permission{
		{Name: "create:course", Description: "Create course"},
		{Name: "delete:course", Description: "Delete course"},
		{Name: "update:course", Description: "Update course"},
		{Name: "view:course", Description: "View course"},
	}

	// Step 2: Create or find existing permissions
	for _, perm := range allPermissions {
		var existing models.Permission
		err := db.Where("name = ?", perm.Name).FirstOrCreate(&existing, perm).Error
		if err != nil {
			return err
		}
	}

	// Step 3: Map name → permission
	permMap := make(map[string]models.Permission)
	for _, perm := range allPermissions {
		var p models.Permission
		db.Where("name = ?", perm.Name).First(&p)
		permMap[perm.Name] = p
	}

	// Step 4: Define roles using shared permission references
	roles := []models.Role{
		{
			Name:        "admin",
			Description: "Administrator role",
			Permissions: []models.Permission{
				permMap["create:course"],
				permMap["delete:course"],
				permMap["update:course"],
				permMap["view:course"],
			},
		},
		{
			Name:        "user",
			Description: "Regular user role",
			Permissions: []models.Permission{
				permMap["view:course"],
			},
		},
	}

	// Step 5: Create roles if not exist
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
