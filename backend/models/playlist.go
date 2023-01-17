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
	StreamingService string
	URI              string
	ImageURL         string
	// This is the foreign key for the tracks that belong to the playlist (one to many relationship)
	Tracks []Track `gorm:"foreignKey:PlaylistID"`
}

// Playlist Response is the model for a JSON response of a playlist
type PlaylistResponse struct {
	Name             string          `json:"name"`
	StreamingService string          `json:"streamingService"`
	URI              string          `json:"uri"`
	ImageURL         string          `json:"imageURL"`
	Tracks           []TrackResponse `json:"tracks"`
}
