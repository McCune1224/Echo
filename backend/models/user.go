package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"not null unique"`
	Email    string `gorm:"not null unique"`
	Password string `gorm:"not null"`
	// This is the foreign key for the playlists that belong to the user (one to many relationship)
	Playlists []Playlist `gorm:"foreignKey:UserID"`
}

type UserResponse struct {
	ID        uint               `json:"id"`
	Username  string             `json:"username"`
	Email     string             `json:"email"`
	Playlists []PlaylistResponse `json:"playlists"`
}
