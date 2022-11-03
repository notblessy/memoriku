package repository

import (
	"github.com/notblessy/memoriku/model"
	"github.com/notblessy/memoriku/utils"
	log "github.com/sirupsen/logrus"
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
	logger := log.WithFields(log.Fields{
		"email": email,
	})

	var user model.User

	err := u.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return &user, nil
}

// FindByID :nodoc:
func (u *userRepository) FindByID(id int64) (*model.User, error) {
	logger := log.WithFields(log.Fields{
		"id": id,
	})

	var user model.User

	err := u.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	user.HidePassword()

	return &user, nil
}

// Update :nodoc:
func (u *userRepository) Update(user *model.User) error {
	logger := log.WithFields(log.Fields{
		"user": utils.Encode(user),
	})

	err := u.db.Save(user).Error
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}
