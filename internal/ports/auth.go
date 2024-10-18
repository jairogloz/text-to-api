package ports

import (
	"context"
	"text-to-api/internal/domain"
)

// AuthService exposes methods for authenticating requests to the API.
type AuthService interface {
	Auth(ctx context.Context, authParams domain.AuthParams) (*domain.RequestContext, error)
}
