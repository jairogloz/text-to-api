package request_context

import (
	"github.com/gofiber/fiber/v2"
	"text-to-api/internal/domain"
	"text-to-api/internal/handlers/ports"
)

type FiberContextKey int

// handler implements the ports.RequestContextHandler interface.
type handler struct {
}

// NewRequestContextHandler returns a new instance of handler implementing the ports.RequestContextHandler interface.
func NewRequestContextHandler() ports.RequestContextHandler {
	return &handler{}
}

// GetClientID retrieves the client ID from the request context.
// It returns an empty string if the client ID is not found.
func (h handler) GetClientID(c *fiber.Ctx) string {
	reqContext, ok := c.Locals(domain.CtxKeyRequestContext).(domain.RequestContext)
	if !ok {
		return ""
	}
	return reqContext.ClientID
}

// GetEnvironment retrieves the environment from the request context.
// It returns nil if the environment is not found.
func (h handler) GetEnvironment(c *fiber.Ctx) *domain.RequestEnvironment {
	reqContext, ok := c.Locals(domain.CtxKeyRequestContext).(domain.RequestContext)
	if !ok {
		return nil
	}
	return reqContext.Environment
}

// GetUserID retrieves the user ID from the request context.
// It returns an empty string if the user ID is not found.
func (h handler) GetUserID(c *fiber.Ctx) string {
	reqContext, ok := c.Locals(domain.CtxKeyRequestContext).(domain.RequestContext)
	if !ok {
		return ""
	}
	return reqContext.UserID
}

// SetClientID sets the client ID in the request context.
func (h handler) SetClientID(c *fiber.Ctx, clientID string) {
	reqContext, ok := c.Locals(domain.CtxKeyRequestContext).(*domain.RequestContext)
	if ok {
		reqContext.ClientID = clientID
		c.Locals(domain.CtxKeyRequestContext, reqContext)
		return
	}
	c.Locals(domain.CtxKeyRequestContext, &domain.RequestContext{
		ClientID: clientID,
	})
}

// SetEnvironment sets the environment in the request context.
func (h handler) SetEnvironment(c *fiber.Ctx, environment *domain.RequestEnvironment) {
	reqContext, ok := c.Locals(domain.CtxKeyRequestContext).(*domain.RequestContext)
	if ok {
		reqContext.Environment = environment
		c.Locals(domain.CtxKeyRequestContext, reqContext)
		return
	}
	c.Locals(domain.CtxKeyRequestContext, &domain.RequestContext{
		Environment: environment,
	})
}

// SetUserID sets the user ID in the request context.
func (h handler) SetUserID(c *fiber.Ctx, userID string) {
	reqContext, ok := c.Locals(domain.CtxKeyRequestContext).(*domain.RequestContext)
	if ok {
		reqContext.UserID = userID
		c.Locals(domain.CtxKeyRequestContext, reqContext)
		return
	}
	c.Locals(domain.CtxKeyRequestContext, &domain.RequestContext{
		UserID: userID,
	})
}
