package domain

// Ptr is a helper function that returns a pointer to the value passed as argument.
func Ptr[T any](val T) *T {
	return &val
}
