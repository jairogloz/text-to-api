package auth

import (
	"context"
	"fmt"
	"text-to-api/internal/crypto"
	"text-to-api/internal/domain"
)

// AuthWithAPIKey authenticates a user by hashing the provided API key and retrieving the corresponding client.
// It returns an AuthResult containing the client ID and environment if authentication is successful.
//
// Parameters:
//   - ctx: The context for managing request deadlines and cancellation signals.
//   - apiKey: The API key used for authentication.
//
// Returns:
//   - A pointer to a domain.AuthResult containing the client ID and environment.
//   - An error if the client cannot be retrieved or any other issue occurs during authentication.
func (s *service) AuthWithAPIKey(ctx context.Context, apiKey string) (*domain.AuthResult, error) {

	// Hash the apiKey
	apiKeyHash := crypto.Hash(apiKey)

	// Search clients by apiKey
	client, apiKeyObject, err := s.clientRepo.GetByAPIKeyHash(ctx, apiKeyHash)
	if err != nil {
		return nil, fmt.Errorf("error getting client by apiKey: %w", err)
	}

	return &domain.AuthResult{
		ClientID:    client.ID,
		Environment: apiKeyObject.Environment,
	}, nil

}
