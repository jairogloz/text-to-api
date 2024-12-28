package ports

import (
	"context"
	"text-to-api/internal/domain"
)

type SubscriptionService interface {
	// CheckSubscription checks if the client has an active subscription for the given environment. If the
	// environment is sandbox, the subscription is considered active.
	// If the environment is live and the subscription is active, the method returns nil, otherwise it returns an error.
	CheckSubscription(ctx context.Context, clientID string, environment domain.RequestEnvironment) (*domain.Subscription, error)
}
