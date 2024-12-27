package ports

import (
	"context"
	"text-to-api/internal/domain"
)

// APIKeyService exposes the methods for operations related to API keys, like creating a new
// APIKey for a client.
type APIKeyService interface {
	CreateAndRevokePrevious(ctx context.Context, reqEnv domain.RequestEnvironment, userID string) (plainAPIKey string, err error)
}

// APIKeyRepository exposes the methods to interact with the storage layer for API keys.
type APIKeyRepository interface {
	SaveAndRevokePrevious(ctx context.Context, newAPIKey domain.APIKey) error
}
