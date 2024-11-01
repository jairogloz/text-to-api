package stripe

import (
	"fmt"
	"github.com/stripe/stripe-go/v80"
)

// Handler is the receiver type for the handler methods related to Stripe.
type Handler struct {
	CheckoutSuccessURL string
	CheckoutCancelURL  string
}

// NewStripeHandler creates a new Stripe handler with the provided API key and URLs.
func NewStripeHandler(apiKey string, successURL, cancelURL string) (*Handler, error) {
	stripe.Key = apiKey
	h := &Handler{
		CheckoutSuccessURL: successURL,
		CheckoutCancelURL:  cancelURL,
	}
	if h.CheckoutSuccessURL == "" {
		return nil, fmt.Errorf("success URL cannot be empty")
	}
	if h.CheckoutCancelURL == "" {
		return nil, fmt.Errorf("cancel URL cannot be empty")
	}
	return h, nil
}
