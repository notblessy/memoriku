package model

import (
	"gorm.io/gorm"
	"time"
)

// CategoryRepository :nodoc:
type CategoryRepository interface {
	Create(cat Category) error
	Update(cat *Category) (err error)
	FindAll(req CategoryReqQuery) (cat *[]Category, count int64, err error)
	FindByID(id int64) (cat *Category, err error)
	DeleteByID(id int64) error
}

// Category :nodoc:
type Category struct {
	ID        int64          `gorm:"primary_key" json:"id"`
	GroupID   int64          `json:"group_id"`
	Name      string         `json:"name" validate:"required"`
	CreatedAt time.Time      `gorm:"<-:create" json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`
}

// CategoryReqQuery :nodoc:
type CategoryReqQuery struct {
	Name string `json:"name"`
	Size int    `json:"size"`
	Page int    `json:"page"`
}
