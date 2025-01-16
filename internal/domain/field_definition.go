package domain

import "fmt"

// FieldDefinition is a struct that represents a field definition in the user's API.
// It can be a request body field, a query parameter or a URL parameter.
type FieldDefinition struct {
	AllowedValues          []interface{}          `json:"allowed_values" bson:"allowed_values"`
	DependentAllowedValues DependentAllowedValues `json:"dependent_allowed_values" bson:"dependent_allowed_values"`
	DependsOn              *string                `json:"depends_on" bson:"depends_on"`
	Description            string                 `json:"description" bson:"description"`
	Name                   string                 `json:"name" bson:"name"`
	Required               *bool                  `json:"required" bson:"required"`
	Type                   string                 `json:"type" bson:"type"`
}

// Validate checks if the FieldDefinition struct has valid data.
// It returns an error if the Name field is empty, if the Type field is empty,
// or if the DependentAllowedValues slice is empty when the field depends on another field.
func (fd *FieldDefinition) Validate() error {
	if fd.DependsOn != nil {
		if len(fd.DependentAllowedValues) == 0 {
			return fmt.Errorf("%w: dependent allowed values are required when field '%s' depends on another field", ErrorValidation, fd.Name)
		}
	}
	if fd.Name == "" {
		return fmt.Errorf("%w: name is required", ErrorValidation)
	}
	if fd.Type == "" {
		return fmt.Errorf("%w: type is required for field '%s'", ErrorValidation, fd.Name)
	}
	return nil
}
