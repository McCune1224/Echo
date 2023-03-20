package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func GetUser(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "TODO: Implement GetUser",
	})
}
