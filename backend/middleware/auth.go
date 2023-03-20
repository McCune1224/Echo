package middleware

import (
	"github.com/McCune1224/Echo/models"
	"github.com/McCune1224/Echo/repository"
	"github.com/gofiber/fiber/v2"
)

// AuthMiddleware is a middleware that checks if the user is authenticated via a session cookie sent in the request header
// If the user is authenticated, the middleware will continue to the next middleware, if not will 401
func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Check if the user is authenticated
		sessionHeader := c.Get("session")
		if sessionHeader == "" {
			return c.Status(401).JSON(fiber.Map{
				"message": "Unauthorized, session is empty",
			})
		}
		// If the user is authenticated, continue to the next middleware
		dbSession := models.Session{}
		dbResponse := repository.DBConnection.First(dbSession, "session = ?", sessionHeader)
		if dbResponse.Error == nil {
			return c.Status(500).JSON(fiber.Map{
				"message": "Internal Server Error",
			})
		}

		dbUser := models.User{}
		dbResponse = repository.DBConnection.First(dbUser, "id = ?", dbSession.UserID)
		if dbUser.ID == 0 {
			return c.Status(401).JSON(fiber.Map{
				"message": "Bad Request, user not found",
			})
		}

		return c.Next()
	}
}
