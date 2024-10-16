package translations

import (
	"context"
	"encoding/json"
	"fmt"
	"text-to-api/internal/domain"
	"time"
)

// Create creates a new translation based on the user's request.
// It validates the userID and the request, sets the current date, marshals the request to JSON,
// and calls the translator to perform the translation. It returns the created Translation or an error.
func (s service) Create(ctx context.Context, request domain.TranslationRequest, userID string) (*domain.Translation, error) {

	if userID == "" {
		return nil, fmt.Errorf("userID is required")
	}
	if err := request.Validate(); err != nil {
		s.logger.Debug(ctx, "Invalid request", "error", err)
		return nil, fmt.Errorf("invalid request: %w", err)
	}

	now := time.Now().UTC()
	request.CurrentDate = now.Format(time.RFC850)

	requestAsJSON, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("could not marshal request: %w", err)
	}

	mappedObject, err := s.translator.TranslateToObject(ctx, string(requestAsJSON), userID)
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
