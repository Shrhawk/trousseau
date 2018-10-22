package aes

import (
	"testing"

	"github.com/oleiade/trousseau/crypto"
	"github.com/stretchr/testify/assert"
)

var secretServiceTests = map[string]struct {
	in         []byte
	passphrase string
	out        []byte
	err        error
}{
	"non-empty input, with non-empty p-empty input, with non-empty passphrase": {[]byte("bonjour"), "passphrase", []byte("bonjour"), nil},
	"empty input, with non-empty passphrase":                                   {nil, "passphrase", nil, crypto.ErrEmptyPayload},
	"empty input, with empty passphrase":                                       {nil, "", nil, crypto.ErrEmptyPayload},
}

func TestSecretService(t *testing.T) {
	for name, test := range secretServiceTests {
		t.Run(name, func(t *testing.T) {
			ss := NewSecretService(test.passphrase)

			encrypted, err := ss.Encrypt(test.in)
			assert.Equal(t, test.err, err)

			decrypted, err := ss.Decrypt(encrypted)
			assert.Equal(t, test.err, err)
			assert.Equal(t, test.out, decrypted)
		})
	}
}
