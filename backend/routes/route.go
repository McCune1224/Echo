package routes

import (
	"github.com/McCune1224/Echo/handlers"
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

	oAuthRoutes.Get("/spotify", handlers.SpotifyOauth)             // /oauth/spotify
	oAuthRoutes.Get("/spotify/redirect", handlers.SpotifyRedirect) // /oauth/spotify/redirect
}

func UserAuthRoutes(app *fiber.App) {
	userRoutes := app.Group("/users") // /user/

	// GET ROUTES
	userRoutes.Get("/", middleware.AuthMiddleware(), handlers.GetUser) // /user

	// POST ROUTES
	userRoutes.Post("/login", handlers.Login)       // /user/login
	userRoutes.Post("/register", handlers.Register) // /user/register
	userRoutes.Post("/logout", handlers.Logout)     // /user/logout

	// PUT ROUTES
	userRoutes.Put("/update", middleware.AuthMiddleware(), handlers.Update) // /user/update

	// DELETE ROUTES
	userRoutes.Post("/delete", middleware.AuthMiddleware(), handlers.Delete) // /user/delete
}

func PlaylistRoutes(app *fiber.App) {
	playlistRoutes := app.Group("/playlists", middleware.AuthMiddleware())

	// GET ROUTES
	playlistRoutes.Get("/", handlers.GetAllPlaylists) // /playlist
	playlistRoutes.Get("/:id", handlers.GetPlaylist)  // /playlist/:id

	// POST ROUTES
	playlistRoutes.Post("/create", handlers.CreatePlaylist) // /playlist/create

	// PUT ROUTES
	playlistRoutes.Put("/update/:id/track", handlers.UpdatePlaylistTracks) // /playlist/update/:id/tracks
	playlistRoutes.Put("/update/:id", handlers.UpdatePlaylistDetails)      // /playlist/update/:id

	// DELETE ROUTES
	playlistRoutes.Delete("/delete/:id", handlers.DeletePlaylist)             // /playlist/delete/:id
	playlistRoutes.Delete("/delete/:id/track", handlers.DeletePlaylistTracks) // /playlist/delete/:id/tracks
}
