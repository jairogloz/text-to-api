package middleware

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"text-to-api/internal/domain"
	hdlPorts "text-to-api/internal/handlers/ports"
	"text-to-api/internal/ports"
)

// UsageLimitMdlwHdl is the receiver for the usage limit middleware handler.
type UsageLimitMdlwHdl struct {
	logger    ports.Logger
	reqCtxHdl hdlPorts.RequestContextHandler
	service   ports.UsageLimitService
}

// NewUsageLimitMdlwHdl creates a new instance of UsageLimitMdlwHdl.
func NewUsageLimitMdlwHdl(l ports.Logger, reqCtxHdl hdlPorts.RequestContextHandler, s ports.UsageLimitService) (*UsageLimitMdlwHdl, error) {
	h := &UsageLimitMdlwHdl{
		logger:    l,
		reqCtxHdl: reqCtxHdl,
		service:   s,
	}
	if h.logger == nil {
		return nil, fmt.Errorf("non-nil logger is required to create a new instance of UsageLimitMdlwHdl")
	}
	if h.reqCtxHdl == nil {
		return nil, fmt.Errorf("non-nil request context handler is required to create a new instance of UsageLimitMdlwHdl")
	}
	if h.service == nil {
		return nil, fmt.Errorf("non-nil usage limit service is required to create a new instance of UsageLimitMdlwHdl")
	}
	return h, nil
}

// UsageLimit is a middleware that checks if a client has exceeded their allowed number of requests.
// If the client ID or environment is not found in the request context, it returns a 500 Internal Server Error.
// If the usage limit is exceeded, it returns a 429 Too Many Requests error.
// Otherwise, it allows the request to proceed to the next handler.
//
// Returns:
//   - fiber.Handler: The middleware handler function.
func (h *UsageLimitMdlwHdl) UsageLimit() fiber.Handler {
	return func(c *fiber.Ctx) error {
		h.logger.Debug(c.Context(), "Usage limit middleware")

		clientID := h.reqCtxHdl.GetClientID(c)
		if clientID == "" {
			h.logger.Error(c.Context(), "Client ID not found in request context")
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Internal Server Error",
			})
		}
		environment := h.reqCtxHdl.GetEnvironment(c)
		if environment == "" {
			h.logger.Error(c.Context(), "Environment not found in request context")
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Internal Server Error",
			})
		}

		err := h.service.CheckUsageLimit(c.Context(), clientID, environment)
		if err != nil {
			if errors.Is(err, domain.ErrorUsageLimitExceeded) {
				return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
					"error": err.Error(),
				})
			}
			h.logger.Error(c.Context(), "Error checking usage limit", "error", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Internal Server Error",
			})
		}

		return c.Next()
	}
}
