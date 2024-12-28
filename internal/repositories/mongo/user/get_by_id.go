package user

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"text-to-api/internal/domain"
)

// GetByID retrieves a user from the repository using the provided client ID and user ID.
// It returns a pointer to the domain.User if found, or an error if the user is not found
// or if there is an issue with the database query.
//
// Note: For the correct behavior, the `users` collection should have a unique index on the (client_id, id) fields.
//
// Parameters:
//   - ctx: The context for managing request deadlines and cancellation signals.
//   - reqEnv: The request environment (live or sandbox) to determine which collection to query.
//   - clientID: The ID of the client to filter the user.
//   - id: The ID of the user to retrieve.
//
// Returns:
//   - A pointer to the domain.User if found.
//   - An error if the user is not found or if there is a database query error.
func (r repository) GetByID(ctx context.Context, reqEnv domain.RequestEnvironment, clientID string, id string) (*domain.User, error) {
	// We start assuming that environment is live
	coll := r.userCollectionLive
	if reqEnv == domain.RequestEnvironmentSandbox {
		coll = r.userCollectionSandbox
	}

	filter := bson.M{
		"client_id": clientID,
		"id":        id,
	}

	var user domain.User
	err := coll.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, domain.ErrorNotFound
		}
		r.logger.Error(ctx, "error getting user by id", "error", err.Error())
		return nil, fmt.Errorf("error querying user by id: %s", err.Error())
	}

	return &user, nil
}
