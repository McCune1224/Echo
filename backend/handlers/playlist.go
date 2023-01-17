package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func GetAllPlaylists(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "OMOAIWA MOSHINDARU",
	})
}
