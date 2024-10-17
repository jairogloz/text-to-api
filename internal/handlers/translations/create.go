package translations

import (
	"github.com/gofiber/fiber/v2"
	"text-to-api/internal/domain"
	"text-to-api/internal/handlers"
	"time"
)

// TranslationRequestBody contains the expected fields when creating a new translation.
type TranslationRequestBody struct {
	InputText       string                   `json:"input_text" bson:"input_text"`
	Lang            string                   `json:"lang" bson:"lang"`
	TranslationType string                   `json:"translation_type" bson:"translation_type"`
	TargetObject    *domain.ObjectDefinition `json:"target_object" bson:"target_object"`
	UserID          string                   `json:"userID" bson:"userID"`
}

func (h *Handler) Create(c *fiber.Ctx) error {

	// Todo: get user/userID from context

	var requestBody TranslationRequestBody
	// Parse the requestBody body into the struct
	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to parse requestBody body"})
	}

	translationRequest := domain.TranslationRequest{
		InputText:       requestBody.InputText,
		Lang:            requestBody.Lang,
		TranslationType: requestBody.TranslationType,
		TargetObject:    requestBody.TargetObject,
		CurrentDate:     time.Now().UTC().Format(time.RFC850),
	}

	translation, err := h.service.Create(c.Context(), translationRequest, requestBody.UserID)
	if err != nil {
		httpStatusCode, message := handlers.HandleError(err)
		return c.Status(httpStatusCode).JSON(fiber.Map{"error": message})
	}

	return c.JSON(translation)
}
