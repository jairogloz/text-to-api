package ports

import "text-to-api/internal/domain"

// AuthService exposes methods for authenticating requests to the API.
type AuthService interface {
	Auth(authParams domain.AuthParams) (*domain.RequestContext, error)
}
