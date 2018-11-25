package proxy

import "net/http"

type Backend interface {
	Proxy() http.Handler
	Upload() http.Handler
	Download() http.Handler

	WebSocket(w http.ResponseWriter, r *http.Request)
}

type Proxy interface {
	Start() error

	Stop() error
}
