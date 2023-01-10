package handlers

import (
	"log"
	"os"
	"time"

	"github.com/McCune1224/Echo/models"
	"github.com/McCune1224/Echo/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// LoginHandler handles the login process for a user and returns a JSON response holding a JWT token if successful
func LoginHandler(c *fiber.Ctx) error {
	var userPostData models.User
	err := c.BodyParser(&userPostData)
	if err != nil {
		return err
	}
	log.Println(userPostData)

	// Make sure all required fields are present in data(email, and password)
	if userPostData.Email == "" || userPostData.Password == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "Missing required fields",
		})
	}

	// Check if user doesn't exist
	var existingUser models.User
	repository.DBConnection.Where("email = ?", userPostData.Email).First(&existingUser)
	if existingUser.Email == "" {
		return c.JSON(fiber.Map{
			"message": "User does not exist",
		})
	}

	err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(userPostData.Password))
	if err != nil {
		return c.JSON(fiber.Map{
			"message": "Incorrect password",
		})
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"id":    existingUser.ID,
		"name":  existingUser.Name,
		"email": existingUser.Email,

		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// Send the token to the user
	return c.JSON(fiber.Map{
		"token": t,
	})
}

func LogoutHandler(c *fiber.Ctx) error {
	//Overwite current SessionID data to expire at current time
	//Erase Session Info in DB
	//?Redirect to desired endpoint?
	return c.SendString("TODO: LOGOUT")
}

func RegisterHandler(c *fiber.Ctx) error {
	// Unwrap POST Request data
	var userPostData models.User
	err := c.BodyParser(&userPostData)
	if err != nil {
		return err
	}

	// Make sure all required fields are present in data(email, and password)
	if (userPostData.Email == "" || userPostData.Name == "") || userPostData.Password == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "Email and Password are required",
		})
	}
	// Hash Passowrd
	hash_password, err := bcrypt.GenerateFromPassword([]byte(userPostData.Password), 14)
	userPostData.Password = string(hash_password)
	if err != nil {
		return err
	}

	// Check if user already exists
	var existingUser models.User
	repository.DBConnection.Where("email = ?", userPostData.Email).First(&existingUser)
	if existingUser.Email != "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "User already exists",
		})
	}

	// Create User
	key := repository.DBConnection.Create(&userPostData)
	if key.Error != nil {
		return key.Error
	}

	// Validate neccesary info (Username, Password & Email)
	// Unwrap POST Data into GORM User struct
	// Store User Struct in Database
	return c.JSON(fiber.Map{"message": "User Created"})
}

func DeleteHandler(c *fiber.Ctx) error {
	// Unwrap POST Request Data
	// Delete any related User Session and then User
	return c.SendString("TODO: DELETE")
}
