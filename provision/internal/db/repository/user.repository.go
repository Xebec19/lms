package repository

import (
	"github.com/Xebec19/lms/provision/internal/db/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUser(id string) (*models.User, error)
	CreateUser(user *models.User) error
	UpdateUser(user *models.User) error
	DeleteUser(id string) error
}

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) GetUser(id string) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepo) UpdateUser(user *models.User) error {
	return r.db.Save(user).Error
}

func (r *UserRepo) DeleteUser(id string) error {
	return r.db.Delete(&models.User{}, id).Error
}
