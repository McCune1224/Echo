package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func GetUser(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.SendString("Welcome " + name)

	// res := repository.DBConnection.Where("id = ?", claimsUserID).First(&user)
}
