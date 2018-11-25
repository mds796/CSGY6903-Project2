package proxy

import (
	"bytes"
	"github.com/mds796/CSGY6903-Project2/crypto"
	"io/ioutil"
	"log"
	"mime"
	"mime/multipart"
	"net/http"
)

func encryptFiles(proxy func(*http.Request), symmetric crypto.SymmetricCipher) func(r *http.Request) {
	return func(r *http.Request) {
		proxy(r)
		upload, err := encryptMultiPartUpload(r, symmetric)

		if err != nil {
			log.Printf("Unable to encrypt files in request body. %v", err)
		}

		r.ContentLength = int64(len(upload))
		r.Body = ioutil.NopCloser(bytes.NewReader(upload))
	}
}

func encryptMultiPartUpload(r *http.Request, symmetric crypto.SymmetricCipher) ([]byte, error) {
	err := r.ParseMultipartForm(r.ContentLength)
	if err != nil {
		return nil, err
	}

	mediaType, params, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return nil, err
	}

	if mediaType != "multipart/form-data" {
		return nil, ErrNotMultipartFormData
	}

	boundary := params["boundary"]

	buffer := &bytes.Buffer{}
	writer := multipart.NewWriter(buffer)

	err = writer.SetBoundary(boundary)
	if err != nil {
		return nil, err
	}

	for _, fileHeaders := range r.MultipartForm.File {
		for i := range fileHeaders {
			file := fileHeaders[i]
			if err != nil {
				return nil, err
			}

			open, err := file.Open()
			if err != nil {
				return nil, err
			}

			slurp, err := ioutil.ReadAll(open)
			if err != nil {
				return nil, err
			}

			cipherText, err := symmetric.Encrypt(slurp)
			if err != nil {
				return nil, err
			}

			file.Header.Set("Content-Type", "application/octet-stream")

			part, err := writer.CreatePart(file.Header)
			if err != nil {
				return nil, err
			}

			_, err = part.Write(cipherText)
			if err != nil {
				return nil, err
			}
		}
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
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
