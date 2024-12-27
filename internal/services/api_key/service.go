package api_key

import (
	"fmt"
	"text-to-api/internal/ports"
)

// service implements port.APIKeyService and holds all the required components to create and manage API keys.
type service struct {
	logger     ports.Logger
	repository ports.APIKeyRepository
}

// NewAPIKeyService creates a new service implementing ports.APIKeyService.
func NewAPIKeyService(l ports.Logger, r ports.APIKeyRepository) (ports.APIKeyService, error) {
	s := &service{
		logger:     l,
		repository: r,
	}

	if s.logger == nil {
		return nil, fmt.Errorf("logger can't be nil")
	}
	if s.repository == nil {
		return nil, fmt.Errorf("repository can't be nil")
	}

	return s, nil
}
