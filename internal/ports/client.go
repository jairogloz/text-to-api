package ports

import "text-to-api/internal/domain"

// ClientRepository exposes methods for interacting with the client repository.
type ClientRepository interface {
	GetClientByAPIKey(apiKey string) (*domain.Client, error)
}
