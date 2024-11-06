package stripe

import (
	"context"
	"fmt"
	"github.com/stripe/stripe-go/v80"
	"github.com/stripe/stripe-go/v80/checkout/session"
	"text-to-api/internal/domain"
)

func (h apiHandler) CreateCheckoutSession(ctx context.Context, priceId string) (*stripe.CheckoutSession, error) {
	if priceId == "" {
		h.logger.Warn(ctx, "Price ID cannot be empty")
		return nil, fmt.Errorf("%w: price ID cannot be empty", domain.ErrorValidation)
	}

	params := &stripe.CheckoutSessionParams{
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				Price:    stripe.String(priceId),
				Quantity: stripe.Int64(1),
			},
		},
		Mode:       stripe.String(string(stripe.CheckoutSessionModeSubscription)),
		SuccessURL: stripe.String(h.checkoutSuccessURL),
		CancelURL:  stripe.String(h.checkoutCancelURL),
	}

	// todo: Get user from database
	// todo: Check if user has a stripe customer_id
	// todo: If she has, check if user has an active subscription, if so, return an error, if not, create a new session specifying the customer_id

	newSession, err := session.New(params)
	if err != nil {
		h.logger.Error(ctx, "Failed to create new stripe checkout session", "error", err)
		return nil, fmt.Errorf("failed to create new stripe checkout session: %w", err)
	}

	return newSession, nil
}
