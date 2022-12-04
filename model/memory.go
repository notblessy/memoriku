package model

import (
	"gorm.io/gorm"
	"time"
)

// MemoryRepository :nodoc:
type MemoryRepository interface {
	Create(memory *Memory) error
	FindAll(req MemoryReqQuery) (memories *[]Memory, count int64, err error)
	FindByID(id string) (cat *Memory, err error)
	DeleteByID(id string) error
}

// Memory :nodoc:
type Memory struct {
	ID               string             `gorm:"primary_key" json:"id"`
	CategoryID       string             `gorm:"->:false;<-" json:"category_id,omitempty"`
	Category         *Category          `gorm:"<-:false" json:"category,omitempty"`
	Title            string             `json:"title"`
	Body             string             `json:"body"`
	Photo            string             `json:"photo"`
	CreatedAt        time.Time          `gorm:"<-:create" json:"created_at"`
	UpdatedAt        time.Time          `json:"updated_at"`
	DeletedAt        gorm.DeletedAt     `json:"deleted_at,omitempty"`
	MemoryReferences *[]MemoryReference `gorm:"<-:create" json:"memory_references"`
	Tags             *[]Tag             `gorm:"many2many:memory_tags" json:"tags"`
}

// MemoryReqQuery :nodoc:
type MemoryReqQuery struct {
	Title      string `json:"title"`
	CategoryID string `json:"categoryID"`
	Size       int    `json:"size"`
	Page       int    `json:"page"`
}

// Tag :nodoc:
type Tag struct {
	ID        int64          `gorm:"primary_key" json:"id"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `gorm:"<-:create" json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`
}

// MemoryReference :nodoc:
type MemoryReference struct {
	ID        int64          `gorm:"primary_key" json:"id"`
	MemoryID  int64          `json:"memory_id"`
	Title     string         `json:"title"`
	Link      string         `json:"link"`
	CreatedAt time.Time      `gorm:"<-:create" json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`
}
