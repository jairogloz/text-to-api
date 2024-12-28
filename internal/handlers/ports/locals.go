package ports

import (
	"github.com/gofiber/fiber/v2"
	"text-to-api/internal/domain"
)

// RequestContextHandler is an interface that exposes methods to get and set values
// in a RequestContext object stored in a fiber context.
type RequestContextHandler interface {
	GetClient(c *fiber.Ctx) *domain.Client
	GetClientID(c *fiber.Ctx) string
	GetEnvironment(c *fiber.Ctx) domain.RequestEnvironment
	GetSubscription(c *fiber.Ctx) *domain.Subscription
	GetUserID(c *fiber.Ctx) string
	SetClient(c *fiber.Ctx, client *domain.Client)
	SetClientID(c *fiber.Ctx, clientID string)
	SetEnvironment(c *fiber.Ctx, environment domain.RequestEnvironment)
	SetSubscription(c *fiber.Ctx, subscription *domain.Subscription)
	SetUserID(c *fiber.Ctx, userID string)
}
