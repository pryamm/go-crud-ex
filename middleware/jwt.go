package middleware

import (
	"backend-c-payment-monitoring/exception"
	"backend-c-payment-monitoring/model"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	jwtMiddleware "github.com/gofiber/jwt/v2"
)

// JWTProtected func for specify routes group with JWT authentication.
// See: https://github.com/gofiber/jwt
func JWTProtected() func(*fiber.Ctx) error {
	// Create config for JWT authentication middleware.
	config := jwtMiddleware.Config{
		SigningKey:   []byte(os.Getenv("JWT_SECRET")),
		ContextKey:   "jwt",
		ErrorHandler: jwtError,
	}

	return jwtMiddleware.New(config)
}

func jwtError(c *fiber.Ctx, err error) error {
	// Return status 401 and failed authentication error.
	return c.Status(http.StatusUnauthorized).JSON(model.ApiResponse{
		Code:    http.StatusUnauthorized,
		Message: "Something Wrong",
		Error:   exception.NewString("Unauthorized"),
		Data:    nil,
	})
}
