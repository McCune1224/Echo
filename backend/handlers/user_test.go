package handlers

import (
	"net/http"
	"net/url"
	"testing"
)

// TestLoginHandler tests the LoginHandler function
func TestLoginHandler(t *testing.T) {
	t.Run("Test LoginHandler", func(t *testing.T) {
		data := url.Values{
			"email":    {"foobarbaz@gmail.com"},
			"password": {"superSecretPassword@#32!"},
		}

		// Post request to /user/login with email and password in body
		req, err := http.PostForm("http://localhost:42069/user/login", data)
		if err != nil {
			t.Fatal(err)
		}

		// Check if response is 200
		if req.StatusCode != 200 {
			t.Fatalf("Expected status code 200, got %d", req.StatusCode)
		}
	})
}

func TestRegisterHandler(t *testing.T) {
	t.Run("Test RegisterHandler", func(t *testing.T) {
		data := url.Values{
			"email":    {"foobarbaz@gmail.com"},
			"password": {"superSecretPassword@#32!"},
		}

		// Post request to /user/register with email and password in body
		req, err := http.PostForm("http://localhost:42069/user/register", data)
		if err != nil {
			t.Fatal(err)
		}
		// Check if response is 200
		if req.StatusCode != 200 {
			t.Fatalf("Expected status code 200, got %d", req.StatusCode)
		}
	})
}
