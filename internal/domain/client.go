package domain

import "time"

const CollNameClients = "clients"

// A Client represents a client in the system. A client can have multiple
// API keys.
type Client struct {
	CreatedAt      time.Time `json:"created_at" bson:"created_at"`
	Email          string    `json:"email" bson:"email"`
	ID             string    `json:"id" bson:"id"`
	LastSignInAt   time.Time `json:"last_sign_in_at" bson:"last_sign_in_at"`
	Name           string    `json:"name" bson:"name"`
	Phone          *string   `json:"phone" bson:"phone"`
	Providers      []string  `json:"providers" bson:"providers"`
	SubscriptionID *string   `json:"subscription_id" bson:"subscription_id"`

	// Potentially delete this field
	CustomerID *string `json:"customer_id" bson:"customer_id"`
}

// ClientDataUpdate represents the parameters that can be updated on a client's data.
// Client data is stored on a separate table/collection from the rest of the client data.
type ClientDataUpdate struct {
	CustomerID     *string `json:"customer_id" bson:"customer_id"`
	SubscriptionID *string `json:"subscription_id" bson:"subscription_id"`
}
