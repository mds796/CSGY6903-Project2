package proxy

import (
	"bytes"
	"github.com/mds796/CSGY6903-Project2/crypto"
	"io/ioutil"
	"net/http"
)

func encryptFiles(proxy func(*http.Request), symmetric crypto.SymmetricCipher) func(r *http.Request) {
	return func(r *http.Request) {
		proxy(r)
	}
}

func decryptFiles(symmetric crypto.SymmetricCipher) func(*http.Response) error {
	return func(r *http.Response) error {
		cipherText, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return err
		}

		plaintext, err := symmetric.Decrypt(cipherText)
		if err != nil {
			return err
		}

		r.ContentLength = int64(len(plaintext))
		r.Body = ioutil.NopCloser(bytes.NewReader(plaintext))

		return nil
	}
}
