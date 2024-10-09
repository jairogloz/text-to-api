package translations

import (
	"context"
	"fmt"
	"text-to-api/internal/domain"
)

// Create creates a new translation to the corresponding endpoint based on the user's request.
func (s service) Create(ctx context.Context, request domain.TranslationRequest) (*domain.Translation, error) {

	// Todo: Get the user endpoints from the cache/storage
	availableEndpoints, err := s.endpointsRepository.FindAll(ctx, "user_id")
	if err != nil {
		return nil, fmt.Errorf("could not get user endpoints: %w", err)
	}

	// Todo: Translate the prompt to the corresponding endpoint
	translation, err := s.translator.Translate(ctx, request, availableEndpoints)
	if err != nil {
		return nil, fmt.Errorf("could not translate prompt: %w", err)
	}

	return translation, nil

}
