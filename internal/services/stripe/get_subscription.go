package stripe

import (
	"context"
	"fmt"
	"github.com/stripe/stripe-go/v80"
	"github.com/stripe/stripe-go/v80/subscription"
	"text-to-api/internal/domain"
)

// GetSubscription retrieves a subscription from Stripe and translates it into a domain.Subscription object.
// This method is part of the internal/services/subscription/ports/subscription.go SubscriptionHandler interface,
// not the ports.StripeAPIHandler interface. It is implemented here as a method of the apiHandler struct to
// keep all the Stripe API calls in one place.
//
// Parameters:
//   - ctx: The context for managing request deadlines and cancellation signals.
//   - subscriptionID: The ID of the subscription to retrieve.
//
// Returns:
//   - A pointer to a domain.Subscription object if the subscription is successfully retrieved and translated.
//   - An error if there is any issue retrieving the subscription or translating its status.
func (h apiHandler) GetSubscription(ctx context.Context, subscriptionID string) (*domain.Subscription, error) {
	subs, err := subscription.Get(subscriptionID, nil)
	if err != nil {
		h.logger.Error(ctx, "failed to get subscription", "subscriptionID", subscriptionID, "error", err)
	}

	// Translate the Stripe subscription object to the domain subscription object
	domainSubscription := &domain.Subscription{
		ID: subs.ID,
	}

	// Translate the subscription status
	switch subs.Status {
	case stripe.SubscriptionStatusActive:
		domainSubscription.Status = domain.SubscriptionStatusActive
	case stripe.SubscriptionStatusCanceled:
		domainSubscription.Status = domain.SubscriptionStatusCanceled
	case stripe.SubscriptionStatusIncomplete:
		domainSubscription.Status = domain.SubscriptionStatusIncomplete
	case stripe.SubscriptionStatusIncompleteExpired:
		domainSubscription.Status = domain.SubscriptionStatusIncompleteExpired
	case stripe.SubscriptionStatusPastDue:
		domainSubscription.Status = domain.SubscriptionStatusPastDue
	case stripe.SubscriptionStatusPaused:
		domainSubscription.Status = domain.SubscriptionStatusPaused
	case stripe.SubscriptionStatusTrialing:
		domainSubscription.Status = domain.SubscriptionStatusTrialing
	case stripe.SubscriptionStatusUnpaid:
		domainSubscription.Status = domain.SubscriptionStatusUnpaid
	default:
		h.logger.Error(ctx, "Unknown subscription status", "status", subs.Status, "subscriptionID", subs.ID)
		return nil, fmt.Errorf("unknown subscription status '%s'", subs.Status)
	}

	// Get the subscription product id
	if subs.Items == nil || len(subs.Items.Data) == 0 {
		h.logger.Error(ctx, "No items in subscription", "subscriptionID", subs.ID)
		return nil, fmt.Errorf("no items in subscription '%s'", subs.ID)
	}
	if len(subs.Items.Data) > 1 {
		h.logger.Warn(ctx, "Multiple items in subscription", "subscriptionID", subs.ID)
	}
	domainSubscription.ProductID = subs.Items.Data[0].Price.Product.ID

	return domainSubscription, nil
}
