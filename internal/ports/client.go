package ports

import (
	"context"
	"text-to-api/internal/domain"
)

// ClientRepository exposes methods for interacting with the client repository.
type ClientRepository interface {
	GetClientByAPIKey(ctx context.Context, apiKeyHash string) (*domain.Client, error)
}
