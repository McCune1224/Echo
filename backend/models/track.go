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

// TrackResponse is the model for a JSON response of a track
type TrackResponse struct {
	Name     string `json:"name"`
	Album    string `json:"album"`
	Artist   string `json:"artist"`
	URI      string `json:"uri"`
	ImageURL string `json:"imageURL"`
}

func NewTrackResponse(track *[]Track) *[]TrackResponse {
	// check if track is nil or list of 0, if so return nil
	if track == nil || len(*track) == 0 {
		return nil
	}
	trackResponse := make([]TrackResponse, len(*track))
	for i, t := range *track {
		trackResponse[i] = TrackResponse{
			Name:     t.Name,
			Album:    t.Album,
			Artist:   t.Artist,
			URI:      t.URI,
			ImageURL: t.ImageURL,
		}
	}
	return &trackResponse
}
