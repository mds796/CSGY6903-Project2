package crypto

type SymmetricCipher interface {
	Encrypt(data []byte) ([]byte, error)

	Decrypt(data []byte) ([]byte, error)
}
