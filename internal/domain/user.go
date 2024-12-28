package domain

const CollNameUsers = "users"
const CollNameUsersSandbox = "users-sandbox"

// User represents an end-user of the API. A client can have multiple users.
type User struct {
	// The client ID that this user belongs to
	ClientID string `json:"client_id" bson:"client_id"`
	// The ID provided by the client to this particular user
	ID string `json:"id" bson:"id"`
	// The user's metadata. Data specific to different translators can be found here
	// for instance, the thread ID for an OpenAI assistant.
	Metadata UserMetadata `json:"metadata" bson:"metadata"`
}
