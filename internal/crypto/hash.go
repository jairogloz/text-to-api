package crypto

import (
	"encoding/hex"
	"golang.org/x/crypto/sha3"
)

// Hash generates a hash from a plain text input.
func Hash(input string) string {
	buf := []byte(input)
	h := sha3.New256()
	h.Write(buf)
	h.Sum(nil)
	return hex.EncodeToString(h.Sum(nil))
}
