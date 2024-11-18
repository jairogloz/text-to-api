package request_context

import (
	"github.com/gofiber/fiber/v2"
	"text-to-api/internal/domain"
)

// GetSubscription retrieves the Subscription from the RequestContext stored in the fiber context.
// It returns nil if no RequestContext is found.
//
// Parameters:
//   - c: The fiber context containing the request context.
//
// Returns:
//   - A pointer to the Subscription if found, or nil if no RequestContext is found.
func (h handler) GetSubscription(c *fiber.Ctx) *domain.Subscription {
	reqContext, ok := c.Locals(domain.CtxKeyRequestContext).(*domain.RequestContext)
	if !ok {
		return nil
	}
	return reqContext.Subscription
}

// SetSubscription sets the Subscription in the RequestContext stored in the fiber context.
// If a RequestContext already exists, it updates the Subscription field.
// If no RequestContext exists, it creates a new one with the provided Subscription.
//
// Parameters:
//   - c: The fiber context containing the request context.
//   - subscription: The Subscription to be set in the request context.
func (h handler) SetSubscription(c *fiber.Ctx, subscription *domain.Subscription) {
	reqContext, ok := c.Locals(domain.CtxKeyRequestContext).(*domain.RequestContext)
	if ok {
		reqContext.Subscription = subscription
		c.Locals(domain.CtxKeyRequestContext, reqContext)
		return
	}
	c.Locals(domain.CtxKeyRequestContext, &domain.RequestContext{
		Subscription: subscription,
	})
}
