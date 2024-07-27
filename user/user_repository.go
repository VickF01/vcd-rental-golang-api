package user

import (
	"gorm.io/gorm"
)

type Repository interface {
	CreateUser(user User) error
	GetUserByUsername(username string) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateUser(user User) error {
	return r.db.Create(&user).Error
}

func (r *repository) GetUserByUsername(username string) (User, error) {
	var user User
	err := r.db.Where("username = ?", username).First(&user).Error
	return user, err
}
