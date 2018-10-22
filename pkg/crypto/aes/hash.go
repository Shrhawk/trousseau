package aes

import (
	"crypto/sha256"
)

func HashPassphrase(passphrase string) []byte {
	h := sha256.New()
	h.Write([]byte(passphrase))
	return h.Sum(nil)
}
