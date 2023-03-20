package models

import (
	"gorm.io/gorm"
)

// Track is the model for a track in the database
type Track struct {
	gorm.Model
	// This is the foreign key for the playlist that the track belongs to
	PlaylistID uint   `gorm:"not null"`
	Name       string `gorm:"not null"`
	Album      string
	Artist     string `gorm:"not null"`
	URI        string
	ImageURL   string
}
