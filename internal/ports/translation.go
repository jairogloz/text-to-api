package ports

import (
	"context"
	"text-to-api/internal/domain"
)

// Translator exposes methods to translate human language prompts into the corresponding API Endpoints.
type Translator interface {
	TranslateToObject(ctx context.Context, translationRequest domain.TranslationRequest, user *domain.User) (translatedObject interface{}, newUserMetadata domain.UserMetadata, err error)
}

// TranslationService exposes the services provided by this application regarding translations.
type TranslationService interface {
	Create(ctx context.Context, request domain.TranslationRequest, reqCtx domain.RequestContext) (*domain.Translation, error)
}
