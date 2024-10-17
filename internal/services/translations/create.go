package translations

import (
	"context"
	"fmt"
	"text-to-api/internal/domain"
)

// Create creates a new translation based on the user's request.
// It validates the userID and the request, sets the current date, marshals the request to JSON,
// and calls the translator to perform the translation. It returns the created Translation or an error.
func (s service) Create(ctx context.Context, request domain.TranslationRequest, userID string) (*domain.Translation, error) {

	if userID == "" {
		return nil, fmt.Errorf("%w: userID is required", domain.ErrorValidation)
	}
	if err := request.Validate(); err != nil {
		s.logger.Debug(ctx, "Invalid request", "error", err)
		return nil, fmt.Errorf("invalid request: %w", err)
	}

	mappedObject, err := s.translator.TranslateToObject(ctx, request, userID)
	if err != nil {
		return nil, fmt.Errorf("could not translate prompt: %w", err)
	}

	translation := &domain.Translation{
		ID:                 "",
		TranslationRequest: request,
		MappedObject:       mappedObject,
	}

	return translation, nil

}
