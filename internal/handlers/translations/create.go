package translations

import (
	"github.com/gofiber/fiber/v2"
	"text-to-api/internal/domain"
	"text-to-api/internal/handlers"
	"time"
)

// TranslationRequestBody contains the expected fields when creating a new translation.
// @Description Request body for translation creation
// @Tags translations
// @Param input_text body string true "The input text to translate"
// @Param lang body string true "The target language for translation"
// @Param translation_type body string true "The type of translation"
// @Param target_object body domain.ObjectDefinition false "The target object for the translation"
// @Produce json
type TranslationRequestBody struct {
	InputText       string                   `json:"input_text" bson:"input_text"`
	Lang            string                   `json:"lang" bson:"lang"`
	TranslationType string                   `json:"translation_type" bson:"translation_type"`
	TargetObject    *domain.ObjectDefinition `json:"target_object" bson:"target_object"`
}

// Create creates a new translation
// @Summary Create a new translation
// @Description Create a new translation based on the provided input text, language, and translation type
// @Tags translations
// @Accept  json
// @Produce  json
// @Param request body TranslationRequestBody true "Request body for creating a new translation"
// @Success 200 {object} domain.Translation "Returns the created translation"
// @Failure 400 {object} fiber.Map "Invalid request body"
// @Failure 500 {object} fiber.Map "Internal server error"
// @Router /v1/translations [post]
func (h *Handler) Create(c *fiber.Ctx) error {

	start := time.Now().UTC()

	requestContext, err := handlers.GetRequestContext(c)
	if err != nil {
		h.logger.Error(c.Context(), "Failed to get request context, or request context is invalid", "error", err.Error())
		httpStatusCode, message := handlers.ToHTTPError(err)
		return c.Status(httpStatusCode).JSON(fiber.Map{"error": message})
	}

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

	translation, err := h.service.Create(c.Context(), translationRequest, *requestContext)
	if err != nil {
		httpStatusCode, message := handlers.ToHTTPError(err)
		return c.Status(httpStatusCode).JSON(fiber.Map{"error": message})
	}

	h.logger.Debug(c.Context(), "Translation created successfully", "total_duration", time.Since(start).String())
	return c.JSON(translation)
}
