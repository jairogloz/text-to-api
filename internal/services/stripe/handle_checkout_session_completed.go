package stripe

import (
	"context"
	"fmt"
	"github.com/stripe/stripe-go/v80"
	"text-to-api/internal/domain"
)

// HandleCheckoutSessionCompleted processes a Stripe Checkout session completion event.
// It validates the session status, extracts the client ID and customer ID from the session,
// and updates the corresponding client data in the database.
//
// Parameters:
//   - ctx: The context for managing request-scoped values, deadlines, and cancellation signals.
//   - session: The Stripe CheckoutSession object containing details about the completed session.
//
// Returns:
//   - error: An error if the session is invalid, required fields are missing, or the database update fails.
func (h apiHandler) HandleCheckoutSessionCompleted(ctx context.Context, session stripe.CheckoutSession) error {

	if session.Status != stripe.CheckoutSessionStatusComplete {
		h.logger.Warn(ctx, "Session is not completed", "status", session.Status)
		return fmt.Errorf("session is not completed")
	}

	clientID := session.ClientReferenceID
	if clientID == "" {
		h.logger.Error(ctx, "session.ClientReferenceID is empty, it's supposed to be the client ID")
		return fmt.Errorf("session.ClientReferenceID is empty, it's supposed to be the client ID")
	}
	if session.Customer == nil {
		h.logger.Error(ctx, "Missing customer ID in session")
		return fmt.Errorf("missing customer ID in session")
	}

	// Update client data
	updateFields := &domain.ClientDataUpdate{
		CustomerID: &session.Customer.ID,
	}
	err := h.clientRepo.UpsertClientData(ctx, clientID, updateFields)
	if err != nil {
		return fmt.Errorf("failed to update client data: %w", err)
	}

	return nil
}
