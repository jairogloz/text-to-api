package ports

import "text-to-api/internal/domain"

// AuthHandler exposes methods for authenticating requests to the API.
type AuthHandler interface {
	Auth(authParams domain.AuthParams) (*domain.RequestContext, error)
}
