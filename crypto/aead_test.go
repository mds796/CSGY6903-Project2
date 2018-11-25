package crypto

import (
	"encoding/hex"
	"testing"
)

func TestAuthenticatedEncryptionCipher_Encrypt(t *testing.T) {
	c := NewSymmetric(NewPassPhraseKey("test"))
	cipherText, err := c.Encrypt([]byte("Hello, World!"))
	if err != nil {
		t.Error(err)
		return
	}

	secondCipherText, err := c.Encrypt([]byte("Hello, World!"))
	if err != nil {
		t.Error(err)
		return
	}

	firstEncryption := hex.EncodeToString(cipherText)
	secondEncryption := hex.EncodeToString(secondCipherText)
	if firstEncryption == secondEncryption {
		t.Errorf(
			"The first cipher text must not match the second because of the nonce. '%v' != '%v'", firstEncryption, secondEncryption)
	}
}

func TestAuthenticatedEncryptionCipher_Decrypt(t *testing.T) {
	expected := "Hello, World!"

	c := NewSymmetric(NewPassPhraseKey("test"))
	cipherText, err := c.Encrypt([]byte(expected))
	if err != nil {
		t.Error(err)
		return
	}

	actual, err := c.Decrypt(cipherText)
	if err != nil {
		t.Error(err)
		return
	}

	if expected != string(actual) {
		t.Errorf(
			"The decrypted plaitnext text does not match the original message. '%v' != '%v'", actual, expected)
	}
}

func TestAuthenticatedEncryptionCipher_DecryptEdited(t *testing.T) {
	expected := "Hello, World!"

	c := NewSymmetric(NewPassPhraseKey("test"))
	cipherText, err := c.Encrypt([]byte(expected))
	if err != nil {
		t.Error(err)
		return
	}

	cipherText[1], cipherText[2] = cipherText[2], cipherText[1]
	cipherText[3], cipherText[4] = cipherText[4], cipherText[3]

	_, err = c.Decrypt(cipherText)
	if err == nil {
		t.Error("Expected error due to authentication failure.")
		return
	}
}

func TestAuthenticatedEncryptionCipher_DecryptDifferentPassPhrase(t *testing.T) {
	expected := "Hello, World!"

	c := NewSymmetric(NewPassPhraseKey("test"))
	cipherText, err := c.Encrypt([]byte(expected))
	if err != nil {
		t.Error(err)
		return
	}

	c2 := NewSymmetric(NewPassPhraseKey("test2"))
	_, err = c2.Decrypt(cipherText)
	if err == nil {
		t.Error("Expected error due to authentication failure.")
		return
	}
}
