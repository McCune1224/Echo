package oauth

import (
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

// SpotifyOauth is the handler for the spotify oauth route
func SpotifyOauth(c *fiber.Ctx) error {
	spotifyOauth := spotifyConfig()
	url := spotifyOauth.AuthCodeURL("state")

	return c.Redirect(url)
}

func SpotifyRedirect(c *fiber.Ctx) error {
	token, err := spotifyConfig().Exchange(oauth2.NoContext, c.Query("code"))
	if err != nil {
		return c.SendStatus(500)
	}

	if len(token.AccessToken) == 0 {
		return c.SendStatus(500)
	}

	// return c.JSON(token)
	// log.Println(token.AccessToken)
	return c.Redirect(os.Getenv("FRONTEND_DOMAIN") + "?token_spotify=" + token.AccessToken)
}
