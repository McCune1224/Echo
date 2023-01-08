package main

import (
	"os"

	"github.com/McCune1224/Echo/repository"
	"github.com/McCune1224/Echo/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

func main() {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		// Frontend domains
		AllowOrigins: "http://localhost:3000, https://echo-frontend.up.railway.app/",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// Grouping routes
	routes.ThirdPartyOauthRoutes(app)
	routes.UserRoutes(app)

	// Connect to Database
	repository.InitDB(os.Getenv("DATABASE_URL"))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello, Echo!",
		})
	})

	app.Listen(getPort())
}
