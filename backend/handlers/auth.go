package handlers

import (
	"regexp"

	"github.com/McCune1224/Echo/models"
	"github.com/McCune1224/Echo/repository"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *fiber.Ctx) error {
	userData := &models.User{}
	if err := c.BodyParser(userData); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid data",
		})
	}

	// Validate email and username
	if userData.Email == "" && userData.Username == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "Email or Username is required",
		})
	}
	if userData.Password == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "Password Required",
		})
	}

	dbUser := &models.User{}
	dbResult := repository.DBConnection.Where("email = ?", userData.Email).First(&dbUser)
	if dbResult.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Error finding user",
			"error":   dbResult.Error,
		})
	}

	if dbUser.ID == 0 {
		return c.Status(400).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	hashErr := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(userData.Password))
	fart, _ := bcrypt.GenerateFromPassword([]byte(userData.Password), bcrypt.DefaultCost)

	if hashErr != nil {
		return c.Status(400).JSON(fiber.Map{
			"message":  "Invalid password",
			"error":    hashErr.Error(),
			"dbUser":   dbUser.Password,
			"userData": string(fart),
		})
	}
	return c.JSON(fiber.Map{
		"message": "TODO: Implement Return of Session",
	})
}

func Logout(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "TODO: Implement Logout",
	})
}

// handles the registration of a new user
func Register(c *fiber.Ctx) error {
	userData := &models.User{}
	if err := c.BodyParser(userData); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid data",
		})
	}

	// Validate that email is not taken
	existingUser := &models.User{}
	repository.DBConnection.Where("email = ? or username = ?", userData.Email, userData.Username).First(&existingUser)
	if existingUser.ID != 0 {
		return c.Status(400).JSON(fiber.Map{
			"message": "Email already taken",
		})
	}
	if existingUser.Username != "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "Username already taken",
		})
	}

	// validate email format
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !emailRegex.MatchString(userData.Email) {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid email",
		})
	}

	// Validate password length
	if len(userData.Password) < 8 || len(userData.Password) > 32 {
		return c.Status(400).JSON(fiber.Map{
			"message": "Password must be between 8 and 32 characters",
		})
	}

	// Validate password contains number
	if !regexp.MustCompile(`[0-9]`).MatchString(userData.Password) {
		return c.Status(400).JSON(fiber.Map{
			"message": "Password must contain at least one number",
		})
	}

	// Validate password contains special character
	if !regexp.MustCompile(`[!@#$%^&*]`).MatchString(userData.Password) {
		return c.Status(400).JSON(fiber.Map{
			"message": "Password must contain at least one special character",
		})
	}

	// Validate password has upper and lower case
	if !regexp.MustCompile(`[a-z]`).MatchString(userData.Password) || !regexp.MustCompile(`[A-Z]`).MatchString(userData.Password) {
		return c.Status(400).JSON(fiber.Map{
			"message": "Password must contain at least one upper and lower case letter",
		})
	}

	// dbResult := repository.DBConnection.Create(newUser)
	// create new user
	dbResult := repository.DBConnection.Create(userData)
	if dbResult.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Error creating user",
			"error":   dbResult.Error.Error(),
		})
	}

	// return new user id
	return c.JSON(fiber.Map{
		"message": "User created",
		"id":      userData.ID,
	})
}

func Update(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "TODO: Implement Update",
	})
}

func Delete(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "TODO: Implement Delete",
	})
}
