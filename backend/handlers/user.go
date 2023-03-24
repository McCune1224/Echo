package handlers

import (
	"github.com/McCune1224/Echo/models"
	"github.com/McCune1224/Echo/repository"
	"github.com/gofiber/fiber/v2"
)

func GetUser(c *fiber.Ctx) error {
	session := c.Locals("session").(*models.Session)

	user := &models.User{}
	dbResult := repository.DBConnection.Where("id = ?", session.UserID).First(&user)
	if dbResult.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Error getting user",
			"error":   dbResult.Error,
		})
	}
	return c.JSON(
		fiber.Map{
			"username": user.Username,
			"email":    user.Email,
			"id":       user.ID,
		},
	)
}
