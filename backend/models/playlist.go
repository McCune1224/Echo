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

// Playlist Response is the model for a JSON response of a playlist
type PlaylistResponse struct {
	ID               uint            `json:"id"`
	Name             string          `json:"name"`
	Description      string          `json:"description"`
	StreamingService string          `json:"streamingService"`
	URI              string          `json:"uri"`
	ImageURL         string          `json:"imageURL"`
	Tracks           []TrackResponse `json:"tracks"`
}

func NewPlaylistResponse(retPlaylist *Playlist) *PlaylistResponse {
	// check if playlist is nil, if so return nil
	if retPlaylist == nil {
		return nil
	}
	// check if playlist has no tracks, if so return playlist without tracks
	if len(retPlaylist.Tracks) == 0 {
		return &PlaylistResponse{
			ID:               retPlaylist.ID,
			Name:             retPlaylist.Name,
			Description:      retPlaylist.Description,
			StreamingService: retPlaylist.StreamingService,
			URI:              retPlaylist.URI,
			ImageURL:         retPlaylist.ImageURL,
		}
	}
	return &PlaylistResponse{
		ID:               retPlaylist.ID,
		Name:             retPlaylist.Name,
		Description:      retPlaylist.Description,
		StreamingService: retPlaylist.StreamingService,
		URI:              retPlaylist.URI,
		ImageURL:         retPlaylist.ImageURL,
		Tracks:           *NewTrackResponse(&retPlaylist.Tracks),
	}
}
