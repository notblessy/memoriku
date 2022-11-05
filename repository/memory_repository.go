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

// Create :nodoc:
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

// Update :nodoc:
func (m *memoryRepository) Update(memory *model.Memory) error {
	logger := log.WithFields(log.Fields{
		"memory": utils.Encode(memory),
	})

	err := m.db.Save(memory).Error
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

// FindAll :nodoc:
func (m *memoryRepository) FindAll(req model.MemoryReqQuery) (memories *[]model.Memory, count int64, err error) {
	logger := log.WithFields(log.Fields{
		"memoryRequest": req,
	})

	offset := (req.Page - 1) * req.Size

	err = m.db.Model(memories).
		Count(&count).
		Error
	if err != nil {
		logger.Error(err)
		return memories, count, err
	}

	err = m.db.Model(memories).
		Limit(req.Size).
		Offset(offset).
		Order("created_at DESC").
		Find(&memories).Error

	if err != nil {
		logger.Error(err)
		return nil, int64(0), err
	}

	return memories, count, err
}

// FindByID :nodoc:
func (m *memoryRepository) FindByID(id int64) (memories *model.Memory, err error) {
	logger := log.WithFields(log.Fields{
		"memoryID": id,
	})

	err = m.db.Take(&memories, id).Error
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return memories, err
}

// DeleteByID :nodoc:
func (m *memoryRepository) DeleteByID(id int64) error {
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
