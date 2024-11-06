package ports

import (
	"context"
	"github.com/stripe/stripe-go/v80"
)

// StripeAPIHandler exposes methods to interact with the Stripe API.
type StripeAPIHandler interface {
	CreateCheckoutSession(ctx context.Context, priceId string) (*stripe.CheckoutSession, error)
}
