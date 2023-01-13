package handlers

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func UserHandler(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	// check if authHeader is empty
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized, no token provided",
		})
	}
	authHeader = authHeader[7:]

	// Put authHeader into a JWT token and then parse claims
	jwtToken, err := jwt.Parse(authHeader, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"message": "Unauthorized, cannot parse token",
		})
	}

	// Check if token is invalid
	if !jwtToken.Valid {
		return c.Status(401).JSON(fiber.Map{
			"message": "Unauthorized, invalid token",
		})
	}

	claims := jwtToken.Claims.(jwt.MapClaims)

	return c.JSON(fiber.Map{
		"message": "User data",
		"claims":  claims,
	})
}
