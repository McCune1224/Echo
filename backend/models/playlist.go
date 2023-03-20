package models

import (
	"gorm.io/gorm"
)

// Playlist is the model for a playlist in the database
type Playlist struct {
	gorm.Model
	// This is the foreign key for the user that created the playlist
	UserID           uint   `gorm:"not null"`
	Name             string `gorm:"not null"`
	Description      string
	StreamingService string
	URI              string
	ImageURL         string
	// This is the foreign key for the tracks that belong to the playlist (one to many relationship)
	Tracks []Track `gorm:"foreignKey:PlaylistID"`
}
