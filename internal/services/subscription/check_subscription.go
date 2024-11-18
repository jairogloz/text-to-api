package subscription

import (
	"context"
	"fmt"
	"text-to-api/internal/domain"
)

// CheckSubscription checks the subscription status of a client based on the provided client ID and environment.
// It retrieves the client and subscription details, and returns the subscription if it is active or trialing.
// If the environment is sandbox, it returns nil without performing any checks.
//
// Parameters:
//   - ctx: The context for managing request deadlines and cancellation signals.
//   - clientID: The ID of the client whose subscription is being checked.
//   - environment: The environment in which the request is being made.
//
// Returns:
//   - A pointer to a domain.Subscription object if the subscription is active or trialing.
//   - An error if there is any issue retrieving the client or subscription, or if the subscription status is not active or trialing.
func (s service) CheckSubscription(ctx context.Context, clientID string, environment domain.RequestEnvironment) (*domain.Subscription, error) {
	if environment == domain.RequestEnvironmentSandbox {
		return nil, nil
	}

	client, err := s.clientRepository.GetByID(ctx, clientID)
	if err != nil {
		return nil, fmt.Errorf("failed to get client while checking subsciption: %w", err)
	}

	if client.SubscriptionID == nil {
		s.logger.Debug(ctx, "client has no subscription", "clientID", clientID)
		return nil, domain.ErrorSubscriptionCanceled
	}

	subscription, err := s.subscriptionHdl.GetSubscription(ctx, *client.SubscriptionID)
	if err != nil {
		return nil, fmt.Errorf("failed to get subscription while checking subscription: %w", err)
	}

	switch subscription.Status {
	case domain.SubscriptionStatusActive,
		domain.SubscriptionStatusTrialing:
		return subscription, nil
	case domain.SubscriptionStatusCanceled:
		return nil, domain.ErrorSubscriptionCanceled
	case domain.SubscriptionStatusIncomplete:
		return nil, domain.ErrorSubscriptionIncomplete
	case domain.SubscriptionStatusIncompleteExpired:
		return nil, domain.ErrorSubscriptionIncompleteExpired
	case domain.SubscriptionStatusPastDue:
		return nil, domain.ErrorSubscriptionPastDue
	case domain.SubscriptionStatusPaused:
		return nil, domain.ErrorSubscriptionPaused
	case domain.SubscriptionStatusUnpaid:
		return nil, domain.ErrorSubscriptionUnpaid
	default:
		s.logger.Error(ctx, "Unknown subscription status", "status", subscription.Status, "subscriptionID", subscription.ID)
		return nil, fmt.Errorf("unknown subscription status '%s'", subscription.Status)
	}
}
