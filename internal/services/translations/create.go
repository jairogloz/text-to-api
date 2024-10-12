package translations

import (
	"context"
	"encoding/json"
	"fmt"
	"text-to-api/internal/domain"
)

// Create creates a new translation to the corresponding endpoint based on the user's request.
func (s service) Create(ctx context.Context, request domain.TranslationRequest, userID string) (*domain.Translation, error) {

	if userID == "" {
		return nil, fmt.Errorf("userID is required")
	}

	requestAsJSON, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("could not marshal request: %w", err)
	}

	// Todo: Translate the prompt to the corresponding endpoint
	translation, err := s.translator.Translate(ctx, string(requestAsJSON), userID)
	if err != nil {
		return nil, fmt.Errorf("could not translate prompt: %w", err)
	}

	return translation, nil

}
