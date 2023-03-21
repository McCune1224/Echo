package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/McCune1224/Echo/models"
	"github.com/McCune1224/Echo/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

// Adding this for the sake of repeatbility in unit testing
func Setup() *fiber.App {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		// Frontend domains whitelist to allow to pass credentials
		AllowOrigins:     "*",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
	}))

	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to Database
	repository.InitDB(os.Getenv("DATABASE_URL"))

	return app
}

var TEST_APP = Setup()

func TestRegister(t *testing.T) {
	TEST_APP.Post("/users/register", Register)
	type testData struct {
		Input           models.User
		ExpectedStatus  int
		ExpectedMessage string
	}

	foo := "bar"
	fmt.Println(foo)

	t.Run("Register with valid data", func(t *testing.T) {
		validPostData := testData{
			Input: models.User{
				Username: "GoodBoy",
				Email:    "goodboy@gmail.com",
				Password: "SecretLongPassword321!",
			},
			ExpectedStatus: 200,
		}

		// Delete user before and after test incase of repeatbility
		repository.DBConnection.Where("email = ?", validPostData.Input.Email).Delete(&models.User{})

		payloadBuffer := bytes.Buffer{}
		err := json.NewEncoder(&payloadBuffer).Encode(validPostData.Input)
		if err != nil {
			t.Error(err)
		}
		req := httptest.NewRequest("POST", "/users/register", &payloadBuffer)
		req.Header.Set("Content-Type", "application/json")
		res, err := TEST_APP.Test(req)
		if err != nil {
			t.Error(err)
		}

		jsonResponse := map[string]interface{}{}
		err = json.NewDecoder(res.Body).Decode(&jsonResponse)
		if err != nil {
			t.Error(err)
		}

		if res.StatusCode != validPostData.ExpectedStatus {
			t.Errorf("Expected status %d, got %d\n", validPostData.ExpectedStatus, res.StatusCode)
		}

		if jsonResponse["id"] == "" {
			t.Errorf("Expected userID, got %s", jsonResponse["id"])
		}
	})

	t.Run("Register with invalid data", func(t *testing.T) {
		payloads := []testData{
			// Invalid email
			{
				ExpectedStatus: 400,
				Input: models.User{
					Email:    "invalid",
					Username: "ValidUsername",
					Password: "ValidPassword123!",
				},
			},
			// Invalid Username
			{
				ExpectedStatus: 400,
				Input: models.User{
					Email:    "valid@gmail.com",
					Username: "ha",
					Password: "ValidPassword123",
				},
			},
			// Invalid Password (too short)
			{
				ExpectedStatus: 400,
				Input: models.User{
					Email:    "valid@gmail.com",
					Username: "ValidUsername",
					Password: "short",
				},
			},
			// Invalid Password (no numbers)
			{
				ExpectedStatus: 400,
				Input: models.User{
					Email:    "valid@gmail.com",
					Username: "ValidUsername",
					Password: "NoNumbers",
				},
			},
			// Invalid Password (no special characters)
			{
				ExpectedStatus: 400,
				Input: models.User{
					Email:    "valid@gmail.com",
					Username: "ValidUsername",
					Password: "NoSpecialCharacters123",
				},
			},

			// Invalid Password (no special characters)
		}

		for _, payload := range payloads {
			repository.DBConnection.Where("email = ?", payload.Input.Email).Delete(&models.User{})
		}
		for _, payload := range payloads {
			payloadBuffer := bytes.Buffer{}
			err := json.NewEncoder(&payloadBuffer).Encode(payload.Input)
			if err != nil {
				t.Error(err)
			}
			req := httptest.NewRequest("POST", "/users/register", &payloadBuffer)
			req.Header.Set("Content-Type", "application/json")
			res, err := TEST_APP.Test(req)
			if err != nil {
				t.Error(err)
			}
			if res.StatusCode != payload.ExpectedStatus {
				t.Errorf("Expected status %d, got %d", payload.ExpectedStatus, res.StatusCode)
			}
		}

		for _, payload := range payloads {
			repository.DBConnection.Where("email = ?", payload.Input.Email).Delete(&models.User{})
		}
	})
}

func TestLogin(t *testing.T) {
	TEST_APP.Post("/users/login", Login)
	type testData struct {
		Input           models.User
		ExpectedStatus  int
		ExpectedMessage string
	}

	// create targetUser to test with
	targetUser := models.User{
		Username: "testuser",
		Password: "Foobarbaz22!",
		Email:    "tu@gmail.com",
	}
	// Creating user just in case they don't already exist

	t.Run("Login with valid data", func(t *testing.T) {
		validPostData := testData{
			Input:          targetUser,
			ExpectedStatus: 200,
		}

		payloadBuffer := bytes.Buffer{}
		err := json.NewEncoder(&payloadBuffer).Encode(validPostData.Input)
		if err != nil {
			t.Error(err)
		}
		req := httptest.NewRequest("POST", "/users/login", &payloadBuffer)
		req.Header.Set("Content-Type", "application/json")
		res, err := TEST_APP.Test(req)
		if err != nil {
			t.Error(err)
		}

		jsonResponse := map[string]interface{}{}
		err = json.NewDecoder(res.Body).Decode(&jsonResponse)
		if res.StatusCode != validPostData.ExpectedStatus {
			t.Errorf("Expected status %d, got %d\nMessage: %v", validPostData.ExpectedStatus, res.StatusCode, jsonResponse["message"])
		}
	})
}
