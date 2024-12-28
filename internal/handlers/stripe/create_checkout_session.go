package stripe

import (
	"github.com/gofiber/fiber/v2"
	"text-to-api/internal/domain"
	"text-to-api/internal/handlers"
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

	// Try to get clientID from context
	clientID := h.ReqCtxHandler.GetClientID(c)
	if clientID == "" {
		h.Logger.Error(c.Context(), "Failed to get client ID from request context")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": domain.ErrorMessageInternalServerError})
	}

	// Check if user has an active customer_id and subscription already
	newSession, err := h.StripeHandler.CreateCheckoutSession(c.Context(), clientID, requestBody.PriceID)
	if err != nil {
		httpStatusCode, message := handlers.ToHTTPError(err)
		return c.Status(httpStatusCode).JSON(fiber.Map{"error": message})
	}

	// Respond with the newSession ID
	return c.JSON(fiber.Map{"id": newSession.ID})
}
