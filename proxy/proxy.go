package proxy

import "net/http"

type Backend interface {
	Proxy(w http.ResponseWriter, r *http.Request)

	WebSocket(w http.ResponseWriter, r *http.Request)

	Upload(w http.ResponseWriter, r *http.Request)

	Download(w http.ResponseWriter, r *http.Request)
}

type Proxy interface {
	Start() error

	Stop() error
}
