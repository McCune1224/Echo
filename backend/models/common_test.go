package models

import (
	"testing"
)

func TestOverwriteModel(t *testing.T) {
	t.Run("Test for Playlist Struct", func(t *testing.T) {
		playlist := &Playlist{
			Name:        "oldName",
			Description: "oldDescription",
		}
		newPlaylist := &Playlist{
			Name:        "newName",
			Description: "newDescription",
		}
		OverwriteModel(playlist, newPlaylist)
		if playlist.Name != "newName" {
			t.Errorf("Expected playlist name to be test2, got %v", playlist.Name)
		}
		if playlist.Description != "newDescription" {
			t.Errorf("Expected playlist description to be test2, got %v", playlist.Description)
		}
	})
}
