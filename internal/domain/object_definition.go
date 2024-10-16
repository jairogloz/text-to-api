package domain

import "fmt"

// ObjectDefinition represents a definition of an object in the system.
type ObjectDefinition struct {
	Fields []FieldDefinition `json:"fields" bson:"fields"`
}

// Validate validates the object definition.
func (od *ObjectDefinition) Validate() error {
	if len(od.Fields) == 0 {
		return fmt.Errorf("%w: at least one field is required", ErrorValidation)
	}

	for _, field := range od.Fields {
		if err := field.Validate(); err != nil {
			return err
		}
	}

	return nil
}
