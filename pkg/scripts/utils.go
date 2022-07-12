package scripts

import (
	"crypto/sha256"
	"encoding/hex"
)

func Hash(password string) string {
	hash := sha256.Sum256([]byte(password))

	return hex.EncodeToString(hash[:])
}
