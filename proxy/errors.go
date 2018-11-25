package proxy

import "errors"

var ErrNotMultipartFormData = errors.New("content type is not multipart form data")
