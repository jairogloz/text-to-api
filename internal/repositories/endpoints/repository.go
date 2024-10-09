package endpoints

import (
	"text-to-api/internal/ports"
)

// repository implements ports.EndpointRepository and holds all the required
// components to interact with the storage of the available endpoints.
type repository struct {
}

// NewEndpointRepository creates a new repository implementing ports.EndpointRepository.
func NewEndpointRepository() (ports.EndpointRepository, error) {
	r := &repository{}

	return r, nil
}
