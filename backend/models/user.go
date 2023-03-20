package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"not null unique"`
	Email    string `gorm:"not null unique"`
	Password string `gorm:"not null"`
	// This is the foreign key for the playlists that belong to the user (one to many relationship)
	Playlists []Playlist `gorm:"foreignKey:UserID"`
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
