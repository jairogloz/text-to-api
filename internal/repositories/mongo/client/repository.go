package client

import (
	"go.mongodb.org/mongo-driver/mongo"
	"text-to-api/internal/domain"
	"text-to-api/internal/ports"
	"time"
)

// repository implements the ports.ClientRepository interface and holds the required components
// to interact with the client repository.
type repository struct {
	clientCollection *mongo.Collection
	logger           ports.Logger
	timeout          time.Duration
}

// NewClientRepository returns a new instance of repository implementing the ports.ClientRepository interface.
func NewClientRepository(mongoClient *mongo.Client, log ports.Logger, dbName string) (ports.ClientRepository, error) {
	clientCollection := mongoClient.Database(dbName).Collection(domain.CollNameClients)
	r := &repository{
		clientCollection: clientCollection,
		logger:           log,
		timeout:          time.Second * 60,
	}
	return r, nil
}
