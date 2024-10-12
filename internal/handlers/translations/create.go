package translations

import (
	"github.com/gofiber/fiber/v2"
	"text-to-api/internal/domain"
)

func (h *Handler) Create(c *fiber.Ctx) error {

	// Todo: get user/userID from context

	var request domain.TranslationRequest
	// Parse the request body into the struct
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to parse request body"})
	}

	translation, err := h.service.Create(c.Context(), request, "userID")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(translation)
}
