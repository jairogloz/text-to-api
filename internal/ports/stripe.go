package ports

import (
	"context"
	"github.com/stripe/stripe-go/v80"
	"text-to-api/internal/domain"
)

// StripeAPIHandler exposes methods to interact with the Stripe API.
type StripeAPIHandler interface {
	CreateCheckoutSession(ctx context.Context, clientId string, priceId string) (*stripe.CheckoutSession, error)
	GetSubscription(ctx context.Context, subscriptionID string) (*domain.Subscription, error)
	HandleCheckoutSessionCompleted(ctx context.Context, session stripe.CheckoutSession) error
	HandleInvoicePaid(ctx context.Context, invoice stripe.Invoice) error
}
