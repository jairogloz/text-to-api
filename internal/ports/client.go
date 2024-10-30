package ports

import (
	"context"
	"text-to-api/internal/domain"
)

// ClientRepository exposes methods for interacting with the client repository.
type ClientRepository interface {
	GetByAPIKeyHash(ctx context.Context, apiKeyHash string) (client *domain.Client, apiKey *domain.APIKey, err error)
}
