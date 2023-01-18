package routes

import (
	"github.com/McCune1224/Echo/handlers"
	oauth "github.com/McCune1224/Echo/handlers/oauth"
	"github.com/McCune1224/Echo/middleware"
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
	oAuthRoutes.Get("/spotify", oauth.SpotifyOauth)
	oAuthRoutes.Get("/spotify/redirect", oauth.SpotifyRedirect)
}

func UserRoutes(app *fiber.App) {
	userRoutes := app.Group("/user")

	// GET ROUTES
	// userRoutes.Get("/", middleware.Protected(), handlers.UserHandler)
	userRoutes.Get("/", middleware.JWTProtected(), handlers.GetUser)

	// POST ROUTES
	userRoutes.Post("/login", handlers.Login)
	userRoutes.Post("/register", handlers.Register)
	userRoutes.Post("/delete", handlers.Delete)
	userRoutes.Post("/logout", handlers.Logout)
}
