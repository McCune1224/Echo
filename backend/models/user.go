package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"not null unique"`
	Email    string `gorm:"not null unique"`
	Password string `gorm:"not null"`
}

type UserResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
