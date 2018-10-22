package aes

import (
	"testing"

	"github.com/oleiade/trousseau/crypto"
	"github.com/stretchr/testify/assert"
)

var encrypterTests = map[string]struct {
	in         []byte
	passphrase string
	err        error
}{
	"non-empty input, with non-empty passphrase": {[]byte("bonjour"), "passphrase", nil},
	"non-empty input, with empty passphrase":     {[]byte("bonjour"), "", crypto.ErrEmptyPassphrase},
	"empty input, with non-empty passphrase":     {nil, "passphrase", crypto.ErrEmptyPayload},
	"empty input, with empty passphrase":         {nil, "", crypto.ErrEmptyPayload},
}

func TestEncrypter(t *testing.T) {
	for name, test := range encrypterTests {
		t.Run(name, func(t *testing.T) {
			e := NewEncrypter(test.passphrase)

			encrypted, err := e.Encrypt(test.in)
			if err == nil {
				assert.NotNil(t, encrypted)
				assert.NotEmpty(t, encrypted)
			}

			assert.Equal(t, test.err, err)
		})
	}
}
