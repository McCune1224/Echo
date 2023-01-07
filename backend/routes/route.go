package routes

import (
	oauth "github.com/McCune1224/Echo/handlers/oauth"
	"github.com/gofiber/fiber/v2"
)

func ThirdPartyOauthRoutes(app *fiber.App) {
	oAuthRoutes := app.Group("/oauth")

	// ------------------ Spotify ------------------

	// /oauth/spotify
	oAuthRoutes.Get("/spotify", oauth.SpotifyOauthHandler)

	// /oauth/spotify/callback
	oAuthRoutes.Get("/spotify/callback", oauth.SpotifyOauthCallbackHandler)

	// ------------------ Youtube ------------------

}
