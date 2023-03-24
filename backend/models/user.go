package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"not null unique" json:"username"`
	Email    string `gorm:"not null unique" json:"email"`
    Password string `gorm:"not null"`
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(hash), err
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	// Hash password
	u.Password, err = hashPassword(u.Password)
	return
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	// Hash password
	u.Password, err = hashPassword(u.Password)
	return
}
