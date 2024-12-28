package ports

import (
	"context"
	"text-to-api/internal/domain"
)

// UserRepository exposes methods for interacting with the user repository.
type UserRepository interface {
	GetByID(ctx context.Context, reqEnv domain.RequestEnvironment, clientID string, id string) (*domain.User, error)
	Insert(ctx context.Context, reqEnv domain.RequestEnvironment, user *domain.User) error
	Update(ctx context.Context, reqEnv domain.RequestEnvironment, user *domain.User) error
}
