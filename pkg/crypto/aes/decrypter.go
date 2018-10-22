package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"

	"github.com/oleiade/trousseau/crypto"
)

type Decrypter struct {
	Passphrase string
}

func (d Decrypter) Decrypt(ciphertext []byte) ([]byte, error) {
	if ciphertext == nil {
		return nil, crypto.ErrEmptyPayload
	} else if d.Passphrase == "" {
		return nil, crypto.ErrEmptyPassphrase
	}

	key, err := d.Key()
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < gcm.NonceSize() {
		return nil, errors.New("malformed ciphertext")
	}

	return gcm.Open(
		nil,
		ciphertext[:gcm.NonceSize()],
		ciphertext[gcm.NonceSize():],
		nil,
	)

}

func (d Decrypter) Key() ([]byte, error) {
	if d.Passphrase == "" {
		return nil, crypto.ErrEmptyPassphrase
	}

	return HashPassphrase(d.Passphrase), nil
}

func NewDecrypter(passphrase string) *Decrypter {
	return &Decrypter{
		Passphrase: passphrase,
	}
}
