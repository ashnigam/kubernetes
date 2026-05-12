package objx

import (
	"encoding/hex"
	"crypto/sha256"
)

// HashWithKey hashes the specified string using the security key
func HashWithKey(data, key string) string {
	d := sha256.Sum256([]byte(data + ":" + key))
	return hex.EncodeToString(d[:])
}
