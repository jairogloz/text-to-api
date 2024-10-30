package auth

import (
	"context"
	"fmt"
	"text-to-api/internal/crypto"
	"text-to-api/internal/domain"
	"text-to-api/internal/ports"
)

// service implements the ports.AuthService interface
type service struct {
	clientRepo ports.ClientRepository
}

// NewAuthService returns a new instance of service implementing the ports.AuthService interface.
func NewAuthService(clientRepo ports.ClientRepository) (ports.AuthService, error) {
	h := &service{
		clientRepo: clientRepo,
	}

	if h.clientRepo == nil {
		return nil, fmt.Errorf("clientRepo is nil")
	}

	return h, nil
}

// Auth authenticates a user by hashing the provided API key and retrieving the corresponding client.
// It returns a RequestContext containing the client ID, user ID, and environment if authentication is successful.
//
// Parameters:
//   - ctx: The context for managing request deadlines and cancellation signals.
//   - authParams: A domain.AuthParams object containing the API key and user ID for authentication.
//
// Returns:
//   - A pointer to a domain.RequestContext containing the client ID, user ID, and environment.
//   - An error if the client cannot be retrieved or any other issue occurs during authentication.
func (h service) Auth(ctx context.Context, authParams domain.AuthParams) (*domain.RequestContext, error) {
	// Hash the apiKey
	apiKeyHash := crypto.Hash(authParams.APIKey)

	// Search clients by apiKey
	client, apiKey, err := h.clientRepo.GetByAPIKeyHash(ctx, apiKeyHash)
	if err != nil {
		return nil, fmt.Errorf("error getting client by apiKey: %w", err)
	}

	// Return the request context
	return &domain.RequestContext{
		ClientID:    client.ID,
		Environment: apiKey.Environment,
		UserID:      authParams.UserID,
	}, nil
}
