package domain

// FieldDefinition is a struct that represents a field definition in the user's API.
// It can be a request body field, a query parameter or a URL parameter.
type FieldDefinition struct {
	Name          string        `json:"name" bson:"name"`
	Type          string        `json:"type" bson:"type"`
	AllowedValues []interface{} `json:"allowed_values" bson:"allowed_values"`
	Required      bool          `json:"required" bson:"required"`
}
