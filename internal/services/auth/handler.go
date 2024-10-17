package auth

import (
	"fmt"
	"text-to-api/internal/crypto"
	"text-to-api/internal/domain"
	"text-to-api/internal/ports"
)

// handler implements the ports.AuthHandler interface
type handler struct {
	clientRepo ports.ClientRepository
}

// NewAuthHandler returns a new instance of handler implementing the ports.AuthHandler interface.
func NewAuthHandler(clientRepo ports.ClientRepository) (ports.AuthHandler, error) {
	h := &handler{
		clientRepo: clientRepo,
	}

	if h.clientRepo == nil {
		return nil, fmt.Errorf("clientRepo is nil")
	}

	return h, nil
}

// Auth
// Todo: document
func (h handler) Auth(authParams domain.AuthParams) (*domain.RequestContext, error) {
	// Hash the apiKey
	apiKeyHash := crypto.Hash(authParams.APIKey)

	// Search clients by apiKey
	client, err := h.clientRepo.GetClientByAPIKey(apiKeyHash)
	if err != nil {
		return nil, fmt.Errorf("error getting client by apiKey: %w", err)
	}

	// Return the request context
	return &domain.RequestContext{
		ClientID: client.ID,
		UserID:   authParams.UserID,
	}, nil
}
