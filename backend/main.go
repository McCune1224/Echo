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
	jwtware "github.com/gofiber/jwt/v3"
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
		// Frontend domains whitelist to allow to pass credentials
		AllowOrigins:     "http://localhost:3000, https://echo-frontend.up.railway.app/",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: true,
	}))

	app.Use(limiter.New(limiter.Config{
		Max:               5,
		Expiration:        30 * time.Second,
		LimiterMiddleware: limiter.SlidingWindow{},
	}))

	// Grouping routes
	routes.ThirdPartyOauthRoutes(app)
	routes.UserRoutes(app)

	signKey := []byte(os.Getenv("JWT_SECRET"))
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: signKey,
	}))
	// Connect to Database
	repository.InitDB(os.Getenv("DATABASE_URL"))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello, Echo!",
		})
	})

	app.Listen(getPort())
}
