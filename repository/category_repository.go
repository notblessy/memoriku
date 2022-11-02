package repository

import (
	"github.com/notblessy/memoriku/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type categoryRepository struct {
	db *gorm.DB
}

// NewCategoryRepository :nodoc:
func NewCategoryRepository(d *gorm.DB) model.CategoryRepository {
	return &categoryRepository{
		db: d,
	}
}

// Create :nodoc:
func (u *categoryRepository) Create(cat model.Category) error {
	logger := log.WithFields(log.Fields{
		"category": cat,
	})

	err := u.db.Create(&cat).Error
	if err != nil {
		logger.Error(err)
		return err
	}

	return err
}
