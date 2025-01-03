package domain

// Endpoint is a struct that represents an endpoint definition in the user's API.
// It responds to the question: "what is the endpoint that must be hit based on such text request?"
type Endpoint struct {
	Name        string            `json:"name" bson:"name"`
	Path        string            `json:"path" bson:"path"`
	Method      string            `json:"method" bson:"method"`
	RequestBody ObjectDefinition  `json:"request_body" bson:"request_body"`
	QueryParams []FieldDefinition `json:"query_params" bson:"query_params"`
	UrlParams   []FieldDefinition `json:"url_params" bson:"url_params"`
}
