package utils

import (
	"encoding/hex"

	"golang.org/x/crypto/sha3"
)

func HashSHA3(data string) string {
	hash := sha3.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}
