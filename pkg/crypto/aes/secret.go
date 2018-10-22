package aes

type SecretService struct {
	Encrypter
	Decrypter
}

func NewSecretService(passphrase string) *SecretService {
	return &SecretService{
		Encrypter: Encrypter{Passphrase: passphrase},
		Decrypter: Decrypter{Passphrase: passphrase},
	}
}
