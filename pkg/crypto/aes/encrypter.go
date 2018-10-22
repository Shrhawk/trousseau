package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"

	"github.com/oleiade/trousseau/crypto"
)

type Encrypter struct {
	Passphrase string
}

func (e Encrypter) Encrypt(plaintext []byte) ([]byte, error) {
	if plaintext == nil {
		return nil, crypto.ErrEmptyPayload
	} else if e.Passphrase == "" {
		return nil, crypto.ErrEmptyPassphrase
	}

	key, err := e.Key()
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

	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, plaintext, nil), nil
}

func (e Encrypter) Key() ([]byte, error) {
	if e.Passphrase == "" {
		return nil, crypto.ErrEmptyPassphrase
	}

	return HashPassphrase(e.Passphrase), nil
}

func NewEncrypter(passphrase string) *Encrypter {
	return &Encrypter{
		Passphrase: passphrase,
	}
}
