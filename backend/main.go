package main

import (
	"os"
	"time"

	"github.com/McCune1224/Echo/repository"
	"github.com/McCune1224/Echo/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/joho/godotenv/autoload"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":42069"
	} else {
		port = ":" + port
	}

	return port
}

// Adding this for the sake of repeatbility in unit testing
func Setup() *fiber.App {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		// Frontend domains whitelist to allow to pass credentials
		AllowOrigins:     "*",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
	}))

	app.Use(limiter.New(limiter.Config{
		Max:               100,
		Expiration:        30 * time.Second,
		LimiterMiddleware: limiter.SlidingWindow{},
	}))
	// Connect to Database
	repository.InitDB(os.Getenv("DATABASE_URL"))
	// Grouping routes
	routes.RootRoutes(app)
	routes.ThirdPartyOauthRoutes(app)
	routes.UserAuthRoutes(app)
	routes.PlaylistRoutes(app)

	return app
}

func main() {
	app := Setup()
	app.Listen(getPort())
}
