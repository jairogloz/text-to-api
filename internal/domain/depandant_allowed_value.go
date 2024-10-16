package domain

import "fmt"

// DependentAllowedValue represents a list of possible allowed values for a field when another field has a specific value.
type DependentAllowedValue struct {
	// The value of the dependency field that will trigger the allowed values.
	Value string `json:"value" bson:"value"`
	// The allowed values for the field when the dependency field has the specified value.
	AllowedValues []interface{} `json:"allowed_values" bson:"allowed_values"`
}

// Validate checks if the DependentAllowedValue struct has valid data.
// It returns an error if the Value field is empty or if the AllowedValues slice is empty.
func (d DependentAllowedValue) Validate() error {
	if d.Value == "" {
		return fmt.Errorf("%w: value is required", ErrorValidation)
	}
	if len(d.AllowedValues) == 0 {
		return fmt.Errorf("%w: allowed values are required for value '%s'", ErrorValidation, d.Value)
	}
	return nil
}

// DependentAllowedValues represents a list of possible allowed values for a field when another field has a specific value.
type DependentAllowedValues []DependentAllowedValue

// Validate checks if each DependentAllowedValue in the DependentAllowedValues slice has valid data.
// It returns an error if any DependentAllowedValue fails validation.
func (da DependentAllowedValues) Validate() error {
	for _, dependentAllowedValue := range da {
		if err := dependentAllowedValue.Validate(); err != nil {
			return err
		}
	}
	return nil
}
