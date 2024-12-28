package stripe

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/stripe/stripe-go/v80"
	"github.com/stripe/stripe-go/v80/webhook"
)

// StripeWebhook handles Stripe webhook events and processes them based on their type.
//
// @Summary Handle Stripe webhook events
// @Description This endpoint listens to Stripe webhook events and processes specific event types such as `checkout.session.completed`, `invoice.paid`, and `invoice.payment_failed`.
// @Tags Webhooks
// @Accept json
// @Produce json
// @Param Stripe-Signature header string true "Stripe signature for webhook validation"
// @Success 200 {string} string "Event processed successfully"
// @Failure 400 {object} map[string]string "Failed to process event due to bad request"
// @Failure 500 {object} map[string]string "Failed to handle event due to internal server error"
// @Router /stripe/webhook [post]
func (h *Handler) StripeWebhook(c *fiber.Ctx) error {

	event, err := webhook.ConstructEvent(c.Body(), c.Get("Stripe-Signature"), h.WebhookSecret)
	if err != nil {
		h.Logger.Error(c.Context(), "Failed to construct event", "error", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to construct event"})
	}

	switch event.Type {
	case stripe.EventTypeCheckoutSessionCompleted:
		h.Logger.Debug(c.Context(), "Checkout session completed", "session_id", event.ID)

		var session stripe.CheckoutSession
		if err := json.Unmarshal(event.Data.Raw, &session); err != nil {
			h.Logger.Error(c.Context(), "Failed to unmarshal session", "error", err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to unmarshal session"})
		}

		err := h.StripeHandler.HandleCheckoutSessionCompleted(c.Context(), session)
		if err != nil {
			h.Logger.Error(c.Context(), "Failed to handle checkout session completed", "error", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to handle checkout session completed"})
		}
	case stripe.EventTypeInvoicePaid:
		h.Logger.Debug(c.Context(), "Invoice paid", "invoice_id", event.ID)

		var invoice stripe.Invoice
		if err := json.Unmarshal(event.Data.Raw, &invoice); err != nil {
			h.Logger.Error(c.Context(), "Failed to unmarshal invoice", "error", err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to unmarshal invoice"})
		}

		err := h.StripeHandler.HandleInvoicePaid(c.Context(), invoice)
		if err != nil {
			h.Logger.Error(c.Context(), "Failed to handle checkout session completed", "error", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to handle checkout session completed"})
		}

	case stripe.EventTypeInvoicePaymentFailed:
		// Todo: implement service
		h.Logger.Debug(c.Context(), "Invoice payment failed", "invoice_id", event.ID)
	default:
		h.Logger.Debug(c.Context(), "Unhandled event type", "event_type", event.Type)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": fmt.Sprintf("Unhandled event type '%s'", event.Type)})
	}

	return c.SendStatus(fiber.StatusOK)
}
