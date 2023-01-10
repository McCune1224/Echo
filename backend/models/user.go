package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"not null unique"`
	Email    string `gorm:"not null unique"`
	Password string `gorm:"not null"`
}

type UserSession struct {
	gorm.Model
	UserID    int    `gorm:"not null"`
	LoginTime string `gorm:"not null"`
	LastSeen  string `gorm:"not null"`
}
