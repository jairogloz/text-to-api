package domain

import "time"

const (
	APIKeyStatusActive  = "active"
	APIKeyStatusRevoked = "revoked"
)

// APIKey represents an API key in the system. A Client can have
// multiple API keys.
type APIKey struct {
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	Hash        string             `json:"hash" bson:"hash"`
	Environment RequestEnvironment `json:"environment" bson:"environment"`
	Status      string             `json:"status" bson:"status"`
	UserID      string             `json:"user_id" bson:"user_id"`
}
