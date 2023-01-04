package oauth

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
)

// Helper Function to make a Oauth2 config struct for Spotify
func spotifyConfig() *oauth2.Config {
	conf := &oauth2.Config{
		ClientID:     os.Getenv("SPOTIFY_CLIENT_ID"),
		ClientSecret: os.Getenv("SPOTIFY_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("SPOTIFY_REDIRECT_URI"),
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://accounts.spotify.com/authorize",
			TokenURL: "https://accounts.spotify.com/api/token",
		},
		Scopes: []string{
			"user-library-read",
			"user-read-private",
			"user-top-read",
		},
	}

	return conf
}

// SpotifyOauthHandler is the handler for the spotify oauth route
func SpotifyOauthHandler(c *fiber.Ctx) error {
	spotifyOauth := spotifyConfig()
	fmt.Println("FOOBARBAZ", spotifyOauth.RedirectURL)
	url := spotifyOauth.AuthCodeURL("state")
	return c.Redirect(url)
}

// SpotifyOauthCallbackHandler is the handler for the spotify oauth callback route
func SpotifyOauthCallbackHandler(c *fiber.Ctx) error {
	return c.SendString("Congrats! You've successfully authenticated with Spotify!")
}
