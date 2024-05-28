package adapters

type EncrypterInterface interface {
	Encrypt(text string) (string, error)
	Decrypt(text string) (string, error)
}
