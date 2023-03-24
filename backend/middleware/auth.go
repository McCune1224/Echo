package middleware

import (
	"time"

	"github.com/McCune1224/Echo/models"
	"github.com/McCune1224/Echo/repository"
	"github.com/gofiber/fiber/v2"
)

// Will set the 'session' in the context if the session is valid
// Will return a 401 if the session is not valid
func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get auth header
		requestSessionID := c.Get("Authorization")
		if requestSessionID == "" {
			return c.Status(401).JSON(fiber.Map{
				"message": "Unauthorized, session is empty",
			})
		}

		dbSession := &models.Session{}
		dbResult := repository.DBConnection.Where("session_id = ?", requestSessionID).First(&dbSession)
		if dbResult.Error != nil {
			return c.Status(401).JSON(fiber.Map{
				"message": "Unauthorized, session not found",
				"error":   dbResult.Error,
			})
		}

		if dbSession.ExpiresAt.Before(time.Now()) {
			return c.Status(401).JSON(fiber.Map{
				"message": "Unauthorized, session expired",
			})
		}

		// Set the session in the context
		c.Locals("session", dbSession)

		c.JSON(fiber.Map{
			"message":   "Authorized",
			"DbSession": dbSession,
		})

		return c.Next()
	}
}
