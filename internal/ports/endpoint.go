package ports

import (
	"context"
	"text-to-api/internal/domain"
)

// EndpointRepository exposes methods to interact with the storage of the available endpoints.
type EndpointRepository interface {
	FindAll(ctx context.Context, userID string) ([]domain.Endpoint, error)
}
