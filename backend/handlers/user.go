package handlers

import (
	"github.com/McCune1224/Echo/models"
	"github.com/McCune1224/Echo/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func LoginHandler(c *fiber.Ctx) error {
	var userPostData models.User
	err := c.BodyParser(&userPostData)
	if err != nil {
		return err
	}

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

	// Create SessionID using UUID and store in DB
	sessionID := uuid.New()

	return c.JSON(fiber.Map{
		"message":   "Logged in",
		"sessionID": sessionID,
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
	if userPostData.Email == "" || userPostData.Password == "" {
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
	return c.JSON(userPostData)
}

func DeleteHandler(c *fiber.Ctx) error {
	// Unwrap POST Request Data
	// Delete any related User Session and then User
	return c.SendString("TODO: DELETE")
}
