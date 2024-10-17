package domain

const CollNameClients = "clients"

type ClientEnvironment string

const (
	ClientEnvironmentLive    ClientEnvironment = "live"
	ClientEnvironmentSandbox ClientEnvironment = "sandbox"
)

// A Client represents a client in the system. A client can have multiple
// API keys.
type Client struct {
	ID      string   `json:"id" bson:"id"`
	Name    string   `json:"name" bson:"name"`
	APIKeys []APIKey `json:"api_keys" bson:"api_keys"`
}

// APIKey represents an API key in the system. A Client can have
// multiple API keys.
// Todo: Add management fields like created_at, updated_at, etc.
type APIKey struct {
	Hash        string            `json:"hash" bson:"hash"`
	Environment ClientEnvironment `json:"environment" bson:"environment"`
}
