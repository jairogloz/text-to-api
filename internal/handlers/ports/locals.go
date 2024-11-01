package ports

import (
	"github.com/gofiber/fiber/v2"
	"text-to-api/internal/domain"
)

// RequestContextHandler is an interface that exposes methods to get and set values
// in a RequestContext object stored in a fiber context.
type RequestContextHandler interface {
	GetClientID(c *fiber.Ctx) string
	GetEnvironment(c *fiber.Ctx) *domain.RequestEnvironment
	GetUserID(c *fiber.Ctx) string
	SetClientID(c *fiber.Ctx, clientID string)
	SetEnvironment(c *fiber.Ctx, environment *domain.RequestEnvironment)
	SetUserID(c *fiber.Ctx, userID string)
}
