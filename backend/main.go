package main

import (
	"os"

	"github.com/McCune1224/Echo/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/joho/godotenv/autoload"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	} else {
		port = ":" + port
	}

	return port
}

func main() {
	app := fiber.New()
	app.Use(logger.New())

	// Grouping routes
	routes.ThirdPartyOauthRoutes(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello, Echo!",
		})
	})

	app.Listen(getPort())
}
