package proxy

import (
	"bytes"
	"io"
	"io/ioutil"
)

type ByteBuffer struct {
	Data []byte
}

func (w *ByteBuffer) Reader() (reader io.ReadCloser, size int) {
	return ioutil.NopCloser(bytes.NewReader(w.Data)), len(w.Data)
}

func (w *ByteBuffer) Write(p []byte) (n int, err error) {
	w.Data = append(w.Data, p...)

	return len(p), nil
}

func NewWriter(size int64) *ByteBuffer {
	return &ByteBuffer{Data: make([]byte, 0, size)}
}
