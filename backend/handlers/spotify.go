package handlers

import (
	"fmt"
	"os"

	"github.com/McCune1224/Echo/models"
	"github.com/McCune1224/Echo/repository"
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

func SpotifyOauth(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "TODO: Implement SpotifyOauth after fixing auth",
	})
	userID := 42069
	spotifyOauth := spotifyConfig()
	url := spotifyOauth.AuthCodeURL("state")
	url = fmt.Sprintf("%s&user=%d", url, userID)

	return c.Redirect(url)
}

func SpotifyRedirect(c *fiber.Ctx) error {
	spotifyToken, err := spotifyConfig().Exchange(oauth2.NoContext, c.Query("code"))
	userID := c.Query("user")

	user := models.User{}
	dbCtx := repository.DBConnection.First(&user, userID)
	if dbCtx.Error != nil {
		return c.SendStatus(500)
	}

	if err != nil {
		return c.SendStatus(500)
	}

	if len(spotifyToken.AccessToken) == 0 {
		return c.SendStatus(500)
	}

	// return c.JSON(token)
	// log.Println(token.AccessToken)
	// return c.Redirect(os.Getenv("FRONTEND_DOMAIN") + "?token_spotify=" + token.AccessToken + "&refresh_token_spotify=" + token.RefreshToken)
	return c.Redirect("http://localhost:4000/dashboard" + "?token_spotify=" + spotifyToken.AccessToken + "&refresh_token_spotify=" + spotifyToken.RefreshToken)
}
