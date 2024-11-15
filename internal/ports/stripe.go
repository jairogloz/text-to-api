package ports

import (
	"context"
	"github.com/stripe/stripe-go/v80"
)

// StripeAPIHandler exposes methods to interact with the Stripe API.
type StripeAPIHandler interface {
	CreateCheckoutSession(ctx context.Context, clientId string, priceId string) (*stripe.CheckoutSession, error)
	HandleCheckoutSessionCompleted(ctx context.Context, session stripe.CheckoutSession) error
	HandleInvoicePaid(ctx context.Context, invoice stripe.Invoice) error
}
