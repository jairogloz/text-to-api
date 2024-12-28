package ports

import (
	"context"
	"text-to-api/internal/domain"
)

// AuthService exposes methods for authenticating requests to the API.
type AuthService interface {
	AuthWithAPIKey(ctx context.Context, apiKey string) (*domain.AuthResult, error)
	AuthWithToken(ctx context.Context, token string) (*domain.AuthResult, error)
}
