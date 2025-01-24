package randomizer

import (
	"github.com/google/uuid"
	"text-to-api/internal/ports"
)

// handler implements ports.Randomizer.
type handler struct {
}

// NewRandomizer creates a new handler implementing ports.Randomizer.
func NewRandomizer() (ports.Randomizer, error) {
	return &handler{}, nil
}

// RandomUUID generates a random UUID.
func (h handler) RandomUUID() string {
	return uuid.New().String()
}
