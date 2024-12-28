package stripe

import (
	"fmt"
	"github.com/stripe/stripe-go/v80"
	"text-to-api/internal/ports"
)

// apiHandler implements ports.StripeAPIHandler and holds the required components to perform
// operations that involve business logic related to Stripe but also require interaction
// with our system databases.
type apiHandler struct {
	clientRepo         ports.ClientRepository
	checkoutSuccessURL string
	checkoutCancelURL  string
	logger             ports.Logger
}

// NewStripeAPIHandler creates a new instance of apiHandler. It returns an error if any of the
// required components are nil.
func NewStripeAPIHandler(apiKey, successURL, cancelURL string, logger ports.Logger, clientRepo ports.ClientRepository) (ports.StripeAPIHandler, error) {
	stripe.Key = apiKey
	h := &apiHandler{
		clientRepo:         clientRepo,
		checkoutCancelURL:  cancelURL,
		checkoutSuccessURL: successURL,
		logger:             logger,
	}
	if h.clientRepo == nil {
		return nil, fmt.Errorf("nil client repository")
	}
	if h.checkoutSuccessURL == "" {
		return nil, fmt.Errorf("success URL cannot be empty")
	}
	if h.checkoutCancelURL == "" {
		return nil, fmt.Errorf("cancel URL cannot be empty")
	}
	if h.logger == nil {
		return nil, fmt.Errorf("nil logger")
	}
	return h, nil
}
