package aes

import (
	"testing"

	"github.com/oleiade/trousseau/crypto"
	"github.com/stretchr/testify/assert"
)

var decrypterTests = map[string]struct {
	in            []byte
	outPassphrase string
	out           []byte
	err           error
}{
	"non-empty input, with non-empty passphrase": {
		[]byte{ // "bonjour" encrypted with "passphrase" passphrase
			251, 247, 241, 109, 96, 176, 107, 62,
			31, 28, 113, 110, 86, 148, 14, 27,
			234, 94, 83, 164, 114, 84, 124, 27,
			167, 17, 170, 95, 158, 90, 99, 177,
			216, 211, 125,
		},
		"passphrase",
		[]byte("bonjour"),
		nil,
	},
	"non-empty input, with empty passphrase": {
		[]byte{ // "bonjour" encrypted with "passphrase" passphrase
			251, 247, 241, 109, 96, 176, 107, 62,
			31, 28, 113, 110, 86, 148, 14, 27,
			234, 94, 83, 164, 114, 84, 124, 27,
			167, 17, 170, 95, 158, 90, 99, 177,
			216, 211, 125,
		},
		"",
		nil,
		crypto.ErrEmptyPassphrase,
	},
	"empty input, with non-empty passphrase": {nil, "passphrase", nil, crypto.ErrEmptyPayload},
	"empty input, with empty passphrase":     {nil, "", nil, crypto.ErrEmptyPayload},
}

func TestDecrypter(t *testing.T) {
	for name, test := range decrypterTests {
		t.Run(name, func(t *testing.T) {
			d := NewDecrypter(test.outPassphrase)

			decrypted, err := d.Decrypt(test.in)
			assert.Equal(t, test.err, err)
			if err == nil {
				assert.NotNil(t, decrypted)
				assert.NotEmpty(t, decrypted)
			}

		})
	}
}
