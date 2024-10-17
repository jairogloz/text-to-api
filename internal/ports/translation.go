package ports

import (
	"context"
	"text-to-api/internal/domain"
)

// Translator exposes methods to translate human language prompts into the corresponding API Endpoints.
type Translator interface {
	TranslateToObject(ctx context.Context, translationRequest domain.TranslationRequest, userID string) (interface{}, error)
}

// TranslationService exposes the services provided by this application regarding translations.
type TranslationService interface {
	Create(ctx context.Context, request domain.TranslationRequest, userID string) (*domain.Translation, error)
}
