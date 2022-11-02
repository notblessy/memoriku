package model

import (
	"gorm.io/gorm"
	"time"
)

// CategoryRepository :nodoc:
type CategoryRepository interface {
	Create(cat Category) error
	FindAll(req CategoryRequest) (cat *[]Category, count int64, err error)
}

// Category :nodoc:
type Category struct {
	gorm.Model
	ID        int64  `gorm:"primary_key"`
	Name      string `json:"name" validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

// CategoryRequest :nodoc:
type CategoryRequest struct {
	Name string `json:"name"`
	Size int    `json:"size"`
	Page int    `json:"page"`
}
