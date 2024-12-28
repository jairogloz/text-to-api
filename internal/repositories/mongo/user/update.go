package user

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"text-to-api/internal/domain"
)

// Update modifies an existing user in the repository based on the provided request environment.
// It returns an error if there is an issue with the database update.
//
// Parameters:
//   - ctx: The context for managing request deadlines and cancellation signals.
//   - reqEnv: The request environment (live or sandbox) to determine which collection to update.
//   - user: A pointer to the domain.User containing the updated user information.
//
// Returns:
//   - An error if there is a database update error.
func (r repository) Update(ctx context.Context, reqEnv domain.RequestEnvironment, user *domain.User) error {
	// We start assuming that environment is live
	coll := r.userCollectionLive
	if reqEnv == domain.RequestEnvironmentSandbox {
		coll = r.userCollectionSandbox
	}

	filter := bson.M{
		"client_id": user.ClientID,
		"id":        user.ID,
	}

	update := bson.M{
		"$set": user,
	}

	_, err := coll.UpdateOne(ctx, filter, update)
	if err != nil {
		r.logger.Error(ctx, "error updating user", "error", err.Error())
		return fmt.Errorf("error updating user: %w", err)
	}

	return nil
}
