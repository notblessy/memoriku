package repository

import (
	"errors"
	"github.com/labstack/echo/v4"
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

// FindByEmailAndPassword :nodoc:
func (u *userRepository) FindByEmailAndPassword(c echo.Context, user model.User) (*model.User, error) {
	err := u.db.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, nil
	case err != nil:
		return nil, err
	default:
		return &user, nil
	}
}
