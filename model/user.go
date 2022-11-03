package model

import (
	"gorm.io/gorm"
	"time"
)

// UserRepository :nodoc:
type UserRepository interface {
	FindByEmail(email string) (*User, error)
	FindByID(id int64) (*User, error)
	Update(user *User) error
}

// User :nodoc:
type User struct {
	ID        int64          `json:"id" gorm:"primary_key"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	Password  string         `json:"password,omitempty"`
	Photo     string         `json:"photo"`
	CreatedAt time.Time      `gorm:"<-:create" json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`
}

func (u *User) HidePassword() *User {
	u.Password = ""

	return u
}
