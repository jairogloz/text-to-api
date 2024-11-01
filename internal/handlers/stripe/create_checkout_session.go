package stripe

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stripe/stripe-go/v80"
	"github.com/stripe/stripe-go/v80/checkout/session"
)

// createCheckoutSessionRequest is the request body for the CreateCheckoutSession handler.
type createCheckoutSessionRequest struct {
	PriceID string `json:"price_id"`
}

// CreateCheckoutSession godoc
// @Summary Create a new Stripe checkout session
// @Description Create a new Stripe checkout session for a subscription. This is a webhook that will be accessed by the front-end when we need to create a checkout session.
// @Tags Stripe
// @Accept json
// @Produce json
// @Param requestBody body createCheckoutSessionRequest true "Request body containing the price ID"
// @Success 200 {object} map[string]string{"id": "session_id"}
// @Failure 400 {object} map[string]string{"error": "Failed to parse request body"}
// @Failure 500 {object} map[string]string{"error": "Failed to create new session"}
// @Router /v1/checkout-session [post]
func (h *Handler) CreateCheckoutSession(c *fiber.Ctx) error {
	var requestBody createCheckoutSessionRequest
	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to parse request body"})
	}

	// Create the newSession
	params := &stripe.CheckoutSessionParams{
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				Price:    stripe.String(requestBody.PriceID),
				Quantity: stripe.Int64(1),
			},
		},
		Mode:       stripe.String(string(stripe.CheckoutSessionModeSubscription)),
		SuccessURL: stripe.String(h.CheckoutSuccessURL),
		CancelURL:  stripe.String(h.CheckoutCancelURL),
	}

	newSession, err := session.New(params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create new session"})
	}

	// Respond with the newSession ID
	return c.JSON(fiber.Map{"id": newSession.ID})
}
