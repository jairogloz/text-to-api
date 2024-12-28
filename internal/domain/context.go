package domain

// CtxKey is a specialized type for context keys to avoid collisions with
// other packages.
type CtxKey int

const (
	CtxKeyRequestContext CtxKey = iota
)
