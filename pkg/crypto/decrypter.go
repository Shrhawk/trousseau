package crypto

type Decrypter interface {
	Decrypter(data []byte) ([]byte, error)
}
