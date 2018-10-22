package crypto

type Encrypter interface {
	Encrypter(plaintext []byte) ([]byte, error)
}
