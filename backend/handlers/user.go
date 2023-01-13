package handlers

import (
	"os"

	"github.com/McCune1224/Echo/models"
	"github.com/McCune1224/Echo/repository"
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
			"message": "Unauthorized, Invalid Token",
		})
	}

	// Check if token is invalid
	if !jwtToken.Valid {
		return c.Status(401).JSON(fiber.Map{
			"message": "Unauthorized, invalid token",
		})
	}

	// Claims type that uses the map[string]interface{} for JSON decoding
	claims := jwtToken.Claims.(jwt.MapClaims)
	// JWT should have user info in it, so pull it out and use it to get user from DB
	userClaims := claims["user"].(map[string]interface{})
	if userClaims == nil {
		return c.Status(401).JSON(fiber.Map{
			"message": "Unauthorized, invalid token",
		})
	}

	user := &models.User{}
	// JWT claims are stored as float64, so convert to uint
	claimsUserID := uint(userClaims["id"].(float64))
	// Attempt to get the user from the database via email
	res := repository.DBConnection.Where("id = ?", claimsUserID).First(&user)
	// Failed to get user from database
	// Respond with error
	if res.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Cannot find user",
		})
	}

	userResponse := &models.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
	// Change User to UserResponse
	return c.JSON(userResponse)
}
