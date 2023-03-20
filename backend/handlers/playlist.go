package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func GetAllPlaylists(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "TODO: Implement GetAllPlaylists",
	})
}

func GetPlaylist(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "TODO: Implement GetPlaylist",
	})
}

func CreatePlaylist(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "TODO: Implement CreatePlaylist",
	})
}

func UpdatePlaylistTracks(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "TODO: Implement UpdatePlaylistTracks",
	})
}

func UpdatePlaylistDetails(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "TODO: Implement UpdatePlaylistDetails",
	})
}

func DeletePlaylist(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "TODO: Implement DeletePlaylist",
	})
}

func DeletePlaylistTracks(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "TODO: Implement DeletePlaylistTracks",
	})
}
