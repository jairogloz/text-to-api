package stripe

import (
	"context"
	"fmt"
	"github.com/stripe/stripe-go/v80"
	"text-to-api/internal/domain"
)

// HandleInvoicePaid processes the Stripe `invoice.paid` webhook event.
// It validates the invoice status and ensures the presence of required fields (customer and subscription).
// If valid, it updates the client's subscription ID in the database.
//
// Parameters:
//   - ctx: The context for managing request-scoped values, deadlines, and cancellation signals.
//   - invoice: The Stripe Invoice object containing details about the paid invoice.
//
// Returns:
//   - error: An error if the invoice status is invalid, required fields are missing, or the database update fails.
func (h apiHandler) HandleInvoicePaid(ctx context.Context, invoice stripe.Invoice) error {
	if invoice.Status != stripe.InvoiceStatusPaid {
		h.logger.Warn(ctx, "Invalid invoice status during 'invoice.paid' event", "status", invoice.Status)
		return fmt.Errorf("invalid invoice status during 'invoice.paid' event '%s'", invoice.Status)
	}
	if invoice.Subscription == nil {
		h.logger.Warn(ctx, "Nil Subscription in 'invoice.paid' event")
		return fmt.Errorf("nil subscription in 'invoice.paid' event")
	}
	if invoice.Customer == nil {
		h.logger.Warn(ctx, "Nil Customer in 'invoice.paid' event")
		return fmt.Errorf("nil Customer in 'invoice.paid' event")
	}
	customerID := invoice.Customer.ID

	// Update client.subscription_id in supbase public.users_data table
	updateParams := &domain.ClientDataUpdate{
		SubscriptionID: &invoice.Subscription.ID,
	}
	err := h.clientRepo.SetSubscriptionByCustID(ctx, customerID, updateParams)
	if err != nil {
		return fmt.Errorf("failed to udpate client data: %w", err)
	}

	return nil
}
