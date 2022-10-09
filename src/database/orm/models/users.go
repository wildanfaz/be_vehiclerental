package models

import (
	"time"
)

type User struct {
	UserId       string    `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"user_id"`
	Name         string    `gorm:"not null" json:"name"`
	Email        string    `gorm:"not null" json:"email"`
	Role         string    `json:"role,omitempty"`
	Password     string    `gorm:"not null" json:"password,omitempty"`
	Gender       string    `json:"gender,omitempty"`
	Address      string    `json:"address,omitempty"`
	MobileNumber string    `json:"mobile_number,omitempty"`
	DisplayName  string    `json:"display_name,omitempty"`
	BirthDate    string    `json:"birth_date,omitempty"`
	CreatedAt    time.Time `gorm:"default:current_timestamp" json:"created_at"`
	UpdatedAt    time.Time `gorm:"default:current_timestamp" json:"updated_at"`
}

type Users []User
