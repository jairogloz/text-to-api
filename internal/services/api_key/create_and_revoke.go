package api_key

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"text-to-api/internal/crypto"
	"text-to-api/internal/domain"
	"time"
)

// CreateAndRevokePrevious creates a new API key for the given user and request environment,
// revokes any previous API keys for the user, and returns the plain text of the new API key.
//
// Parameters:
//   - ctx: The context for managing request-scoped values, deadlines, and cancellation signals.
//   - reqEnv: The request environment for which the API key is being created.
//   - userID: The ID of the user for whom the API key is being created.
//
// Returns:
//   - string: The plain text of the newly created API key.
//   - error: An error if any of the operations fail.
func (s service) CreateAndRevokePrevious(ctx context.Context, reqEnv domain.RequestEnvironment, userID string) (string, error) {
	if reqEnv == "" {
		s.logger.Warn(ctx, "empty request environment")
		return "", fmt.Errorf("empty request environment")
	}
	if userID == "" {
		s.logger.Warn(ctx, "empty user ID")
		return "", fmt.Errorf("empty user ID")
	}

	// Create a new UUID
	keyPlain := uuid.New()

	// Hash the UUID
	keyHashed := crypto.Hash(keyPlain.String())

	// Build the new APIKey
	apiKey := domain.APIKey{
		CreatedAt:   time.Now().UTC(),
		Hash:        keyHashed,
		Environment: reqEnv,
		Status:      domain.APIKeyStatusActive,
		UserID:      userID,
	}

	// Save the new APIKey to repository and revoke the previous one
	err := s.repository.SaveAndRevokePrevious(ctx, apiKey)
	if err != nil {
		s.logger.Error(ctx, "failed to save and deactivate previous API key", "error", err)
		return "", fmt.Errorf("failed to save and deactivate previous API key: %w", err)
	}

	// This is the only time we return the plain API key
	return keyPlain.String(), nil
}
