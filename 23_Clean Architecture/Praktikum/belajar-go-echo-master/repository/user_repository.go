package repository

import (
	"belajar-go-echo/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll() ([]model.User, error)
	Save(user *model.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) FindAll() ([]model.User, error) {
	users := make([]model.User, 0)
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) Save(user *model.User) error {
	return r.db.Create(user).Error
}
