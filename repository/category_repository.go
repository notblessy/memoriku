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
func (c *categoryRepository) Create(cat model.Category) error {
	logger := log.WithFields(log.Fields{
		"category": cat,
	})

	err := c.db.Create(&cat).Error
	if err != nil {
		logger.Error(err)
		return err
	}

	return err
}

// Update :nodoc:
func (c *categoryRepository) Update(cat *model.Category) error {
	logger := log.WithFields(log.Fields{
		"category": utils.Encode(cat),
	})

	err := c.db.Save(cat).Error
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

// FindAll :nodoc:
func (c *categoryRepository) FindAll(req model.CategoryReqQuery) (cats *[]model.Category, count int64, err error) {
	logger := log.WithFields(log.Fields{
		"categoryRequest": req,
	})

	offset := (req.Page - 1) * req.Size

	err = c.db.Model(cats).
		Count(&count).
		Error
	if err != nil {
		logger.Error(err)
		return cats, count, err
	}

	err = c.db.Model(cats).
		Limit(req.Size).
		Offset(offset).
		Order("created_at DESC").
		Find(&cats).Error

	if err != nil {
		logger.Error(err)
		return nil, int64(0), err
	}

	return cats, count, err
}

// FindByID :nodoc:
func (c *categoryRepository) FindByID(id string) (cat *model.Category, err error) {
	logger := log.WithFields(log.Fields{
		"categoryID": id,
	})

	err = c.db.Take(&cat, id).Error
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return cat, err
}

// DeleteByID :nodoc:
func (c *categoryRepository) DeleteByID(id string) error {
	logger := log.WithFields(log.Fields{
		"categoryID": id,
	})

	var cat model.Category
	err := c.db.Delete(&cat, id).Error
	if err != nil {
		logger.Error(err)
		return err
	}

	return err
}
