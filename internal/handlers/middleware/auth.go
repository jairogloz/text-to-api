package middleware

import (
	"github.com/gofiber/fiber/v2"
	"strings"
	"text-to-api/internal/domain"
	"text-to-api/internal/ports"
)

// AuthMiddleware returns a middleware that checks the Authorization header for a valid token
// using the given authService. If the token is valid, it sets the request context in the fiber context
// so that it can be used in the handler.
func AuthMiddleware(authService ports.AuthHandler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get the Authorization header
		authHeader := c.Get("Authorization")

		// Check if it contains the Bearer prefix
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized, missing or invalid Authorization header",
			})
		}

		// Extract the token from the header
		token := strings.TrimPrefix(authHeader, "Bearer ")

		// Get the User-Id header if present
		userIDHeader := c.Get("User-Id")

		authParams := domain.AuthParams{
			APIKey: token,
			UserID: userIDHeader,
		}

		requestCtx, err := authService.Auth(authParams)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		// Set the request context in the fiber context
		c.Locals(domain.CtxKeyRequestContext, requestCtx)

		return c.Next()
	}
}
