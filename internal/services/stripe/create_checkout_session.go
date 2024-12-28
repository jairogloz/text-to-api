package stripe

import (
	"context"
	"errors"
	"fmt"
	"github.com/stripe/stripe-go/v80"
	"github.com/stripe/stripe-go/v80/checkout/session"
	"text-to-api/internal/domain"
)

// CreateCheckoutSession initializes a new Stripe Checkout session for a given client and price ID.
// This function verifies that a valid price ID is provided and sets up the checkout session parameters
// accordingly. It first retrieves the client data from the database, and if the client already has a
// Stripe customer ID, it attaches this ID to the session. If not, a new Stripe customer ID is generated
// during the session creation.
//
// If the session is created successfully but the client lacks a Stripe customer ID, this function updates
// the client's record with the new customer ID from Stripe.
//
// Parameters:
// - ctx: Context for handling request-scoped values and cancellation signals.
// - clientId: The unique identifier of the client within the application’s database.
// - priceId: The unique identifier for the Stripe price used to set up the subscription.
//
// Returns:
// - *stripe.CheckoutSession: A pointer to the created Stripe Checkout session if successful.
// - error: An error if the session creation fails due to missing parameters, database retrieval issues, or Stripe API errors.
func (h apiHandler) CreateCheckoutSession(ctx context.Context, clientId string, priceId string) (*stripe.CheckoutSession, error) {
	if priceId == "" {
		h.logger.Warn(ctx, "Price ID cannot be empty")
		return nil, fmt.Errorf("%w: price ID cannot be empty", domain.ErrorValidation)
	}
	if clientId == "" {
		h.logger.Warn(ctx, "Client ID cannot be empty")
		return nil, fmt.Errorf("%w: client ID cannot be empty", domain.ErrorValidation)
	}

	params := &stripe.CheckoutSessionParams{
		CancelURL:         stripe.String(h.checkoutCancelURL),
		ClientReferenceID: stripe.String(clientId), // This is the supabase user_id, will be useful to link the stripe customer_id to the user when the session is completed
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				Price:    stripe.String(priceId),
				Quantity: stripe.Int64(1),
			},
		},
		Mode:       stripe.String(string(stripe.CheckoutSessionModeSubscription)),
		SuccessURL: stripe.String(h.checkoutSuccessURL),
	}

	// Get client(user) from database
	client, err := h.clientRepo.GetByID(ctx, clientId)
	if err != nil && !errors.Is(err, domain.ErrorNotFound) {
		h.logger.Error(ctx, "Failed to get client", "error", err)
		return nil, fmt.Errorf("failed to get client: %w", err)
	}
	// Check if client has a stripe customer_id
	if client != nil && client.CustomerID != nil {
		// Get stripe customer
		// todo: Issue #1 If she has, check if client has an active subscription, if so, return an error, if not, create a new session specifying the customer_id
		params.Customer = stripe.String(*client.CustomerID)
	}

	newSession, err := session.New(params)
	if err != nil {
		h.logger.Error(ctx, "Failed to create new stripe checkout session", "error", err)
		return nil, fmt.Errorf("failed to create new stripe checkout session: %w", err)
	}

	return newSession, nil
}
