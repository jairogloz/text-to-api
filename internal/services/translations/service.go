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
	translator ports.Translator
	userRepo   ports.UserRepository
}

// NewTranslationsService creates a new service implementing ports.TranslationService.
func NewTranslationsService(t ports.Translator, l ports.Logger, ur ports.UserRepository) (ports.TranslationService, error) {
	s := &service{
		logger:     l,
		translator: t,
		userRepo:   ur,
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

	return s, nil
}
