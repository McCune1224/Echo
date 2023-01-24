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
	oAuthRoutes := app.Group("/oauth") // /oauth/

	oAuthRoutes.Get("/spotify", oauth.SpotifyOauth)             // /oauth/spotify
	oAuthRoutes.Get("/spotify/redirect", oauth.SpotifyRedirect) // /oauth/spotify/redirect
}

func UserRoutes(app *fiber.App) {
	userRoutes := app.Group("/user") // /user/

	// GET ROUTES
	// userRoutes.Get("/", middleware.Protected(), handlers.UserHandler)
	userRoutes.Get("/", middleware.JWTProtected(), handlers.GetUser) // /user

	// POST ROUTES
	userRoutes.Post("/login", handlers.Login)       // /user/login
	userRoutes.Post("/register", handlers.Register) // /user/register
	userRoutes.Post("/delete", handlers.Delete)     // /user/delete
	userRoutes.Post("/logout", handlers.Logout)     // /user/logout
}

func PlaylistRoutes(app *fiber.App) {
	playlistRoutes := app.Group("/playlist", middleware.JWTProtected())

	// GET ROUTES
	playlistRoutes.Get("/", handlers.GetAllPlaylists) // /playlist
	playlistRoutes.Get("/:id", handlers.GetPlaylist)  // /playlist/:id

	// POST ROUTES
	playlistRoutes.Post("/create", handlers.CreatePlaylist) // /playlist/create

	// PUT ROUTES
	playlistRoutes.Put("/update/:id/tracks", handlers.UpdatePlaylistTracks) // /playlist/update/:id/tracks
	playlistRoutes.Put("/update/:id", handlers.UpdatePlaylistDetails)       // /playlist/update/:id

	// DELETE ROUTES
	playlistRoutes.Delete("/delete/:id", handlers.DeletePlaylist)              // /playlist/delete/:id
	playlistRoutes.Delete("/delete/:id/tracks", handlers.DeletePlaylistTracks) // /playlist/delete/:id/tracks
}
