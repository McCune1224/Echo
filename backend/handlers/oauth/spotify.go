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
	url := spotifyOauth.AuthCodeURL("state")

	// append service name to url
	url = url + "&service=spotify"
	fmt.Print(fmt.Sprintf("Redirecting to %s", url))
	return c.Redirect(url)
}
