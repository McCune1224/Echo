package models

import (
	"testing"
)

func TestNewPlaylistResponse(t *testing.T) {
	t.Run("nil playlist", func(t *testing.T) {
		playlist := NewPlaylistResponse(nil)
		if playlist != nil {
			t.Errorf("Expected nil, got %v", playlist)
		}
	})
	t.Run("valid playlist", func(t *testing.T) {
		playlist := NewPlaylistResponse(&Playlist{
			Name:        "test",
			Description: "test",
		})
		if playlist == nil {
			t.Errorf("Expected playlist, got nil")
		}
		if playlist.Name != "test" {
			t.Errorf("Expected playlist name to be test, got %v", playlist.Name)
		}
		if playlist.Description != "test" {
			t.Errorf("Expected playlist description to be test, got %v", playlist.Description)
		}
	})
}
