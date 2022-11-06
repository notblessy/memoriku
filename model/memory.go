package model

import (
	"gorm.io/gorm"
	"time"
)

// MemoryRepository :nodoc:
type MemoryRepository interface {
	Upsert(memory *Memory) error
	FindAll(req MemoryReqQuery) (memories *[]Memory, count int64, err error)
	FindByID(id int64) (cat *Memory, err error)
	DeleteByID(id int64) error
}

// Memory :nodoc:
type Memory struct {
	ID               int64              `gorm:"primary_key" json:"id"`
	CategoryID       int64              `gorm:"->:false;<-" json:"category_id,omitempty"`
	Category         Category           `json:"category"`
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
	Name string `json:"name"`
	Size int    `json:"size"`
	Page int    `json:"page"`
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
