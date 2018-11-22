package proxy

import "net/http"

type Server interface {
	Proxy(w http.ResponseWriter, r *http.Request)

	Upload(w http.ResponseWriter, r *http.Request)

	Download(w http.ResponseWriter, r *http.Request)
}
