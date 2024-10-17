package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"text-to-api/internal/domain"
)

// GetRequestContext returns the request context from the fiber context.
// If the request context is not found, it returns an error.
func GetRequestContext(c *fiber.Ctx) (*domain.RequestContext, error) {
	requestContext, ok := c.Locals(domain.CtxKeyRequestContext).(*domain.RequestContext)
	if !ok {
		return nil, fmt.Errorf("request context not found in fiber context")
	}
	err := requestContext.Validate()
	if err != nil {
		return nil, err
	}
	return requestContext, nil
}
