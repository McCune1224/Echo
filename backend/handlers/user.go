package handlers

import (
	"github.com/McCune1224/Echo/models"
	"github.com/McCune1224/Echo/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func GetUser(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	claimsUserID := claims["id"]
	// Convert claimsUserID to uint from float64
	claimsUserID = uint(claimsUserID.(float64))

	var dbUser models.User
	// Query the database for the user with the id
	res := repository.DBConnection.Where("id = ?", claimsUserID).First(&dbUser)
	if res.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Error getting user",
		})
	}

	// Convert dbUser to UserResponse
	userResponse := models.UserResponse{
		ID:       dbUser.ID,
		Username: dbUser.Username,
		Email:    dbUser.Email,
	}

	// Create a new user object to return to the client

	return c.JSON(userResponse)
}
