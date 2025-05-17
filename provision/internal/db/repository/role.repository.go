package repository

import (
	"github.com/Xebec19/lms/provision/internal/db/models"
	"gorm.io/gorm"
)

type RoleRepository interface {
	GetRole(name string) (*models.Role, error)
}

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{
		db: db,
	}
}

func (r *roleRepository) GetRole(name string) (*models.Role, error) {
	var role models.Role
	if err := r.db.Preload("Permissions").Where("name = ?", name).First(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}
