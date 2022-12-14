package repository

import (
	"github.com/notblessy/memoriku/model"
	"github.com/notblessy/memoriku/utils"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type memoryRepository struct {
	db *gorm.DB
}

// NewMemoryRepository :nodoc:
func NewMemoryRepository(d *gorm.DB) model.MemoryRepository {
	return &memoryRepository{
		db: d,
	}
}

// Upsert :nodoc:
func (m *memoryRepository) Create(memory *model.Memory) error {
	logger := log.WithFields(log.Fields{
		"memory": utils.Encode(memory),
	})

	tx := m.db.Begin()

	err := tx.Create(&memory).Error
	if err != nil {
		logger.Error(err)
		tx.Rollback()
		return err
	}

	tx.Commit()
	return err
}

// FindAll :nodoc:
func (m *memoryRepository) FindAll(req model.MemoryReqQuery) (memories *[]model.Memory, count int64, err error) {
	logger := log.WithFields(log.Fields{
		"memoryRequest": req,
	})

	// Query Builder for filter memories
	qm := m.db.Model(memories)
	qt := m.db.Model(memories)

	offset := (req.Page - 1) * req.Size

	if req.Title != "" {
		qm.Where("title LIKE ?", "%"+req.Title+"%")
		qt.Where("title LIKE ?", "%"+req.Title+"%")
	}

	if req.CategoryID != "" {
		qm.Where("category_id = ?", req.CategoryID)
		qt.Where("category_id = ?", req.CategoryID)
	}

	err = qt.Count(&count).Error
	if err != nil {
		logger.Error(err)
		return memories, count, err
	}

	err = qm.Joins("Category").
		Preload("MemoryReferences").
		Preload("Tags").
		Limit(req.Size).Offset(offset).
		Order("created_at DESC").
		Find(&memories).Error

	if err != nil {
		logger.Error(err)
		return nil, int64(0), err
	}

	return memories, count, err
}

// FindByID :nodoc:
func (m *memoryRepository) FindByID(id string) (memories *model.Memory, err error) {
	logger := log.WithFields(log.Fields{
		"memoryID": id,
	})

	err = m.db.Joins("Category").
		Preload("MemoryReferences").
		Preload("Tags").Take(&memories, id).Error
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return memories, err
}

// DeleteByID :nodoc:
func (m *memoryRepository) DeleteByID(id string) error {
	logger := log.WithFields(log.Fields{
		"memoryID": id,
	})

	var memory model.Memory
	err := m.db.Delete(&memory, id).Error
	if err != nil {
		logger.Error(err)
		return err
	}

	return err
}

func ScopeCategory(categoryID string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if categoryID != "" {
			return db.Where("category_id = ?", categoryID)
		}
		return db
	}
}
