package trousseau

import (
	"github.com/oleiade/trousseau/crypto/openpgp"
)

// Declare encryption types
type CryptoType int
const (
	SYMMETRIC_ENCRYPTION  CryptoType = iota
	ASYMMETRIC_ENCRYPTION
)

// Declare available encryption algorithms
type CryptoAlgorithm int
const (
	GPG_ENCRYPTION 		CryptoAlgorithm = iota
	AES_256_ENCRYPTION
)

func DecryptAsymmetricPGP(encryptedData []byte, passphrase string) ([]byte, error) {
	// Decrypt store data
	decryptionKeys, err := openpgp.ReadSecRing(openpgp.SecringFile)
	if err != nil {
		return nil, err
	}

	plainData, err := openpgp.Decrypt(decryptionKeys, string(encryptedData), passphrase)
	if err != nil {
		return nil, err
	}

	return plainData, nil
}

func EncryptAsymmetricPGP(plainData []byte, recipients []string) ([]byte, error) {
	encryptionKeys, err := openpgp.ReadPubRing(openpgp.PubringFile, recipients)
	if err != nil {
		return nil, err
	}

	encData := openpgp.Encrypt(encryptionKeys, string(plainData))

	return encData, nil
}