package ports

// Randomizer exposes methods for generating random values.
type Randomizer interface {
	RandomUUID() string
}
