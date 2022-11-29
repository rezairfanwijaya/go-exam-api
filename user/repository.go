package user

import (
	"errors"

	"gorm.io/gorm"
)

// interface
type IUserRepository interface {
	FindByEmail(email string) (User, error)
}

type userRepository struct {
	db *gorm.DB
}

// function new repo
func NewRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

// function to handle search user by email
func (r *userRepository) FindByEmail(email string) (User, error) {
	var user User

	if err := r.db.Where("email = ?", email).Find(&user).Error; err != nil {
		return user, errors.New("user not found")
	}

	return user, nil
}
