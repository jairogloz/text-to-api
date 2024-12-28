package api_key

import (
	"github.com/gofiber/fiber/v2"
	"text-to-api/internal/domain"
	"text-to-api/internal/handlers"
)

// Create creates a new API key and revokes the previous one
// @Summary Create a new API key
// @Description Create a new API key for the client and revoke the previous one
// @Tags api_key
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]string "Returns the new API key"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /v1/api_keys [post]
func (h *Handler) Create(c *fiber.Ctx) error {
	// Get clientID from context. Remember that clientID here is the user_id in supabase.
	clientID := h.ReqCtxHandler.GetClientID(c)
	if clientID == "" {
		h.Logger.Error(c.Context(), "clientID not found in context")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": domain.ErrorMessageInternalServerError,
		})
	}
	// Get environment from context
	environment := h.ReqCtxHandler.GetEnvironment(c)
	if environment == "" {
		h.Logger.Error(c.Context(), "environment not found in context")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": domain.ErrorMessageInternalServerError,
		})
	}

	// Consume service to create API key
	plainAPIKey, err := h.Service.CreateAndRevokePrevious(c.Context(), environment, clientID)
	if err != nil {
		h.Logger.Error(c.Context(), "failed to create API key", "error", err)
		httpStatusCode, message := handlers.ToHTTPError(err)
		return c.Status(httpStatusCode).JSON(fiber.Map{"error": message})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"api_key": plainAPIKey,
	})
}
