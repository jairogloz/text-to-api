package user

import (
	"go.mongodb.org/mongo-driver/mongo"
	"text-to-api/internal/domain"
	"text-to-api/internal/ports"
	"time"
)

// repository implements the ports.UserRepository interface and holds the required components
// to interact with the user repository.
type repository struct {
	userCollectionLive    *mongo.Collection
	userCollectionSandbox *mongo.Collection
	logger                ports.Logger
	timeout               time.Duration
}

// NewUserRepository returns a new instance of repository implementing the ports.UserRepository interface.
func NewUserRepository(mongoClient *mongo.Client, log ports.Logger, dbName string) (ports.UserRepository, error) {
	userCollectionLive := mongoClient.Database(dbName).Collection(domain.CollNameUsers)
	r := &repository{
		userCollectionLive: userCollectionLive,
		logger:             log,
		timeout:            time.Second * 60,
	}
	return r, nil
}
