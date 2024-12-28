package request_context

import (
	"github.com/gofiber/fiber/v2"
	"text-to-api/internal/domain"
)

// GetClient retrieves the Client from the RequestContext stored in the fiber context.
// It returns nil if no RequestContext is found.
//
// Parameters:
//   - c: The fiber context containing the request context.
//
// Returns:
//   - A pointer to the Client if found, or nil if no RequestContext is found.
func (h handler) GetClient(c *fiber.Ctx) *domain.Client {
	reqContext, ok := c.Locals(domain.CtxKeyRequestContext).(*domain.RequestContext)
	if !ok {
		return nil
	}
	return reqContext.Client
}

// SetClient sets the Client in the RequestContext stored in the fiber context.
// If a RequestContext already exists, it updates the Client field.
// If no RequestContext exists, it creates a new one with the provided Client.
//
// Parameters:
//   - c: The fiber context containing the request context.
//   - client: The Client to be set in the request context.
func (h handler) SetClient(c *fiber.Ctx, client *domain.Client) {
	reqContext, ok := c.Locals(domain.CtxKeyRequestContext).(*domain.RequestContext)
	if ok {
		reqContext.Client = client
		c.Locals(domain.CtxKeyRequestContext, reqContext)
		return
	}
	c.Locals(domain.CtxKeyRequestContext, &domain.RequestContext{
		Client: client,
	})
}
