package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

// AuthMiddleware is a middleware that checks if the user i s authenticated with a JWT token before allowing access to a route.
func JWTProtected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	})
}
