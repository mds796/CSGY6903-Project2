package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"io"
)

type AuthenticatedEncryptionCipher struct {
	PassPhrase string
}

func (c *AuthenticatedEncryptionCipher) createCipher() (cipher.AEAD, error) {
	block, err := aes.NewCipher(createHash(c.PassPhrase))
	if err != nil {
		return nil, err
	}

	return cipher.NewGCM(block)
}

func (c *AuthenticatedEncryptionCipher) Encrypt(data []byte) ([]byte, error) {
	aead, err := c.createCipher()
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, aead.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return aead.Seal(nonce, nonce, data, nil), nil
}

func (c *AuthenticatedEncryptionCipher) Decrypt(data []byte) ([]byte, error) {
	aead, err := c.createCipher()
	if err != nil {
		return nil, err
	}

	nonceSize := aead.NonceSize()
	nonce, cipherText := data[:nonceSize], data[nonceSize:]

	plaintext, err := aead.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

func createHash(key string) []byte {
	hash := sha256.New()
	hash.Write([]byte(key))

	return hash.Sum(nil)
}

func NewSymmetric(passPhrase string) SymmetricCipher {
	return &AuthenticatedEncryptionCipher{PassPhrase: passPhrase}
}
