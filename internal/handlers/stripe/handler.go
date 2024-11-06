package stripe

import (
	"fmt"
	httpPorts "text-to-api/internal/handlers/ports"
	"text-to-api/internal/ports"
)

// Handler is the receiver type for the handler methods related to Stripe.
type Handler struct {
	Logger        ports.Logger
	ReqCtxHandler httpPorts.RequestContextHandler
	StripeHandler ports.StripeAPIHandler
}

// NewStripeHandler creates a new Stripe handler with the provided API key and URLs.
func NewStripeHandler(logger ports.Logger, stripeHandler ports.StripeAPIHandler, reqCtxHandler httpPorts.RequestContextHandler) (*Handler, error) {
	h := &Handler{
		Logger:        logger,
		ReqCtxHandler: reqCtxHandler,
		StripeHandler: stripeHandler,
	}
	if h.Logger == nil {
		return nil, fmt.Errorf("nil logger")
	}
	if h.ReqCtxHandler == nil {
		return nil, fmt.Errorf("nil request context handler")
	}
	if h.StripeHandler == nil {
		return nil, fmt.Errorf("nil Stripe handler")
	}

	return h, nil
}
