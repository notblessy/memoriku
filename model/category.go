package model

import (
	"gorm.io/gorm"
	"time"
)

var GroupCategory = map[string]string{
	"PROGRAMMING": "Programming",
	"TRAVEL":      "Travel",
}

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
	GroupID   string         `json:"group_id,omitempty"`
	Name      string         `json:"name" validate:"required"`
	CreatedAt time.Time      `gorm:"<-:create" json:"created_at,omitempty"`
	UpdatedAt time.Time      `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`
}

// CategoryWeb :nodoc:
type CategoryWeb struct {
	GroupID    string        `json:"group_id"`
	Categories []ValueObject `json:"categories"`
}

// CategoryReqQuery :nodoc:
type CategoryReqQuery struct {
	Name string `json:"name"`
	Size int    `json:"size"`
	Page int    `json:"page"`
}
