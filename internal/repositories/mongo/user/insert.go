package user

import (
	"context"
	"fmt"
	"text-to-api/internal/domain"
)

// Insert adds a new user to the repository based on the provided request environment.
// It returns an error if there is an issue with the database insertion.
//
// Parameters:
//   - ctx: The context for managing request deadlines and cancellation signals.
//   - reqEnv: The request environment (live or sandbox) to determine which collection to insert into.
//   - user: A pointer to the domain.User to be inserted.
//
// Returns:
//   - An error if there is a database insertion error.
func (r repository) Insert(ctx context.Context, reqEnv domain.RequestEnvironment, user *domain.User) error {
	if user == nil {
		return fmt.Errorf("user is required")
	}

	// We start assuming that environment is live
	coll := r.userCollectionLive
	if reqEnv == domain.RequestEnvironmentSandbox {
		coll = r.userCollectionSandbox
	}

	_, err := coll.InsertOne(ctx, user)
	if err != nil {
		r.logger.Error(ctx, "error inserting user", "error", err.Error())
		return fmt.Errorf("error inserting user: %s", err.Error())
	}

	return nil
}
