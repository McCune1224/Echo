package routes

import (
	"github.com/McCune1224/Echo/handlers"
	oauth "github.com/McCune1224/Echo/handlers/oauth"
	"github.com/gofiber/fiber/v2"
)

func RootRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello, Echo!",
		})
	})
}

func ThirdPartyOauthRoutes(app *fiber.App) {
	oAuthRoutes := app.Group("/oauth")

	// ------------------ Spotify ------------------

	// /oauth/spotify
	oAuthRoutes.Get("/spotify", oauth.SpotifyOauthHandler)

	// /oauth/spotify/callback
	oAuthRoutes.Get("/spotify/callback", oauth.SpotifyOauthCallbackHandler)

	// ------------------ Youtube ------------------
}

func UserRoutes(app *fiber.App) {
	userRoutes := app.Group("/user")

	// GET ROUTES
	// userRoutes.Get("/", middleware.Protected(), handlers.UserHandler)
	userRoutes.Get("/", handlers.UserHandler)

	// POST ROUTES
	userRoutes.Post("/login", handlers.LoginHandler)
	userRoutes.Post("/register", handlers.RegisterHandler)
	userRoutes.Post("/delete", handlers.DeleteHandler)
}
