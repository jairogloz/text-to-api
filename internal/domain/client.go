package domain

import "time"

const CollNameClients = "clients"

// A Client represents a client in the system. A client can have multiple
// API keys.
type Client struct {
	CreatedAt    time.Time `json:"created_at" bson:"created_at"`
	CustomerID   *string   `json:"customer_id" bson:"customer_id"`
	Email        string    `json:"email" bson:"email"`
	ID           string    `json:"id" bson:"id"`
	LastSignInAt time.Time `json:"last_sign_in_at" bson:"last_sign_in_at"`
	Name         string    `json:"name" bson:"name"`
	Phone        *string   `json:"phone" bson:"phone"`
	Providers    []string  `json:"providers" bson:"providers"`
}

// ClientDataUpdate represents the parameters that can be updated on a client's data.
// Client data is stored on a separate table/collection from the rest of the client data.
type ClientDataUpdate struct {
	CustomerID *string `json:"customer_id" bson:"customer_id"`
}

// APIKey represents an API key in the system. A Client can have
// multiple API keys.
// Todo: Add management fields like created_at, updated_at, etc.
type APIKey struct {
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	Hash        string             `json:"hash" bson:"hash"`
	Environment RequestEnvironment `json:"environment" bson:"environment"`
	Status      string             `json:"status" bson:"status"`
	UserID      string             `json:"user_id" bson:"user_id"`
}
