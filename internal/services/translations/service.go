package translations

import (
	"fmt"
	"text-to-api/internal/ports"
)

// service implements ports.TranslationService and holds all the required
// components to translate human language prompts into the corresponding
// API Endpoints.
type service struct {
	logger     ports.Logger
	randomizer ports.Randomizer
	translator ports.Translator
	userRepo   ports.UserRepository
}

// NewTranslationsService creates a new service implementing ports.TranslationService.
func NewTranslationsService(t ports.Translator, l ports.Logger, ur ports.UserRepository, r ports.Randomizer) (ports.TranslationService, error) {
	s := &service{
		logger:     l,
		translator: t,
		userRepo:   ur,
		randomizer: r,
	}

	if s.logger == nil {
		return nil, fmt.Errorf("logger can't be nil")
	}
	if s.translator == nil {
		return nil, fmt.Errorf("translator can't be nil")
	}
	if s.userRepo == nil {
		return nil, fmt.Errorf("user repository can't be nil")
	}
	if s.randomizer == nil {
		return nil, fmt.Errorf("randomizer can't be nil")
	}

	return s, nil
}
