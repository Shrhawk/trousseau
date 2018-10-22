package crypto

type SecretService interface {
	Encrypter
	Decrypter
}
