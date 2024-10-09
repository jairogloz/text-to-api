package translations

import (
	"github.com/gofiber/fiber/v2"
	"text-to-api/internal/domain"
)

func (h *Handler) Create(c *fiber.Ctx) error {
	translation, err := h.service.Create(c.Context(), domain.TranslationRequest{})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(translation)
}
