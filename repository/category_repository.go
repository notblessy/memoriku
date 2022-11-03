package repository

import (
	"github.com/notblessy/memoriku/model"
	"github.com/notblessy/memoriku/utils"
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

// Update :nodoc:
func (u *categoryRepository) Update(cat *model.Category) error {
	logger := log.WithFields(log.Fields{
		"user": utils.Encode(cat),
	})

	err := u.db.Save(cat).Error
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

// FindAll :nodoc:
func (u *categoryRepository) FindAll(req model.CategoryRequest) (cat *[]model.Category, count int64, err error) {
	logger := log.WithFields(log.Fields{
		"categoryRequest": req,
	})

	offset := (req.Page - 1) * req.Size

	err = u.db.Model(cat).
		Count(&count).
		Error
	if err != nil {
		logger.Error(err)
		return cat, count, err
	}

	err = u.db.Model(cat).
		Limit(req.Size).
		Offset(offset).
		Order("created_at DESC").
		Find(&cat).Error

	if err != nil {
		logger.Error(err)
		return nil, int64(0), err
	}

	return cat, count, err
}

// FindByID :nodoc:
func (u *categoryRepository) FindByID(id int64) (cat *model.Category, err error) {
	logger := log.WithFields(log.Fields{
		"categoryID": id,
	})

	err = u.db.Take(&cat, id).Error
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return cat, err
}
