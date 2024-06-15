package model

import (
	"time"
	"gorm.io/gorm"
)

type UserRequest struct {
	Email       string `json:"email" validate:"required,email"`
	PhoneNumber string `json:"phone_number" validate:"omitempty"`
	FirstName   string `json:"first_name" validate:"omitempty"`
	LastName    string `json:"last_name" validate:"omitempty"`
	Role        string `json:"role" validate:"required,oneof=admin user"`
	BirthDay    string `json:"birth_day" validate:"omitempty,datetime=2006-01-02T15:04:05Z07:00"`
}

type User struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Email       string    `json:"email" gorm:"unique;not null"`
	PhoneNumber string    `json:"phone_number" gorm:"unique"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Role        string    `json:"role" gorm:"type:enum('admin','user');default:'user'"`
	BirthDay    time.Time `json:"birth_day"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
}

type Error struct {
	ResponseCode int    `json:"response_code"`
	Message      string `json:"message"`
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&User{})
}