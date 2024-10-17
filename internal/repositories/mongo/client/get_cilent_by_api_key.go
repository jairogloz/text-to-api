package client

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"text-to-api/internal/domain"
)

// GetClientByAPIKey retrieves a client from the repository using the provided API key hash.
// It returns a pointer to the domain.Client if found, or an error if the client is not found
// or if there is an issue with the database query.
//
// Parameters:
//   - ctx: The context for managing request deadlines and cancellation signals.
//   - apiKeyHash: The hash of the API key used to filter the client.
//
// Returns:
//   - A pointer to the domain.Client if found.
//   - An error if the client is not found or if there is a database query error.
func (r repository) GetClientByAPIKey(ctx context.Context, apiKeyHash string) (*domain.Client, error) {
	filter := bson.M{"api_keys.hash": apiKeyHash}

	var client domain.Client
	err := r.clientCollection.FindOne(ctx, filter).Decode(&client)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, domain.ErrorNotFound
		}
		r.logger.Error(ctx, "error getting client by apiKey", "error", err.Error())
		return nil, fmt.Errorf("failed to get client by apiKey: %w", err)
	}

	return &client, nil
}
