package translations

import (
	"fmt"
	"text-to-api/internal/ports"
)

// service implements ports.TranslationService and holds all the required
// components to translate human language prompts into the corresponding
// API Endpoints.
type service struct {
	endpointsRepository ports.EndpointRepository
	translator          ports.Translator
}

// NewTranslationsService creates a new service implementing ports.TranslationService.
func NewTranslationsService(t ports.Translator, endpointsRepo ports.EndpointRepository) (ports.TranslationService, error) {
	s := &service{
		endpointsRepository: endpointsRepo,
		translator:          t,
	}

	if s.endpointsRepository == nil {
		return nil, fmt.Errorf("endpointsRepository can't be nil")
	}
	if s.translator == nil {
		return nil, fmt.Errorf("translator can't be nil")
	}

	return s, nil
}
