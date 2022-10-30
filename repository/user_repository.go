package repository

import (
	"errors"
	"github.com/notblessy/memoriku/model"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

// NewUserRepository :nodoc:
func NewUserRepository(d *gorm.DB) model.UserRepository {
	return &userRepository{
		db: d,
	}
}

// FindByEmail :nodoc:
func (u *userRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	err := u.db.Where("email = ?", email).First(&user).Error
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, nil
	case err != nil:
		return nil, err
	default:
		return &user, nil
	}
}
