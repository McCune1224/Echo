package handlers

import (
	"github.com/McCune1224/Echo/models"
	"github.com/McCune1224/Echo/repository"
	"github.com/gofiber/fiber/v2"
)

func GetAllPlaylists(c *fiber.Ctx) error {
	// Get the userToken from the JWT
	userToken := c.Locals("user").(*models.User)
	db := repository.DBConnection
	var playlists []models.Playlist

	// Get all playlists for the user
	db.Where("user_id = ?", userToken.ID).Find(&playlists)

	// Convert the Playlist struct to PlaylistResponse struct
	var playlistResponses []models.PlaylistResponse
	for _, playlist := range playlists {
		playlistResponses = append(playlistResponses, models.PlaylistResponse{
			Name: playlist.Name,
		})
	}
}
