package model

import (
	"gorm.io/gorm"
	"time"
)

// UserRepository :nodoc:
type UserRepository interface {
	FindByEmail(email string) (*User, error)
}

// User :nodoc:
type User struct {
	gorm.Model
	ID        int64  `gorm:"primary_key"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Photo     string `json:"photo"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
