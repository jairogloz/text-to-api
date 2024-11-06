package stripe

import (
	"fmt"
	"text-to-api/internal/ports"
)

// Handler is the receiver type for the handler methods related to Stripe.
type Handler struct {
	StripeHandler ports.StripeAPIHandler
}

// NewStripeHandler creates a new Stripe handler with the provided API key and URLs.
func NewStripeHandler(stripeHandler ports.StripeAPIHandler) (*Handler, error) {
	h := &Handler{
		StripeHandler: stripeHandler,
	}
	if h.StripeHandler == nil {
		return nil, fmt.Errorf("nil Stripe handler")
	}

	return h, nil
}
