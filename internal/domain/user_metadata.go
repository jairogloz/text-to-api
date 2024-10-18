package domain

// UserMetadata represents the metadata of a user.
type UserMetadata map[string]interface{}

// GetString returns the string value of the key if it exists, otherwise it returns an empty string.
func (um UserMetadata) GetString(key string) string {
	if v, ok := um[key]; ok {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}

// Equals compares the current UserMetadata with another UserMetadata to determine if they are equal.
// It returns true if both UserMetadata have the same length and identical key-value pairs.
//
// Parameters:
//   - newUserMetadata: The UserMetadata to compare with the current UserMetadata.
//
// Returns:
//   - A boolean value indicating whether the two UserMetadata are equal.
func (um UserMetadata) Equals(newUserMetadata UserMetadata) bool {
	if len(um) != len(newUserMetadata) {
		return false
	}

	for k, v := range um {
		if newUserMetadata[k] != v {
			return false
		}
	}

	return true
}
