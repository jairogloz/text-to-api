package ports

import (
	"context"
	"text-to-api/internal/domain"
)

// StripeSubscriptionHandler exposes methods to interact with Stripe subscriptions.
type StripeSubscriptionHandler interface {
	GetSubscription(ctx context.Context, subscriptionID string) (*domain.Subscription, error)
}
