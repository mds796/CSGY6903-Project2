package proxy

import (
	"log"
	"net/http"
)

type HttpServerProxy struct {
	Multiplexer *http.ServeMux
	Server      *http.Server
	Backend     Backend
}

func (p *HttpServerProxy) Start() error {
	p.configureRoutes()
	return p.listenAndServe()
}

func (p *HttpServerProxy) configureRoutes() {
	p.Multiplexer.HandleFunc("/", p.Backend.Proxy)
	p.Multiplexer.HandleFunc("/upload", p.Backend.Upload)
	p.Multiplexer.HandleFunc("/download", p.Backend.Download)
}

func (p *HttpServerProxy) listenAndServe() error {
	log.Printf("Proxy server now listening on %v.\n", p.Server.Addr)

	return p.Server.ListenAndServe()
}

func (p *HttpServerProxy) Stop() error {
	return p.Server.Close()
}

func NewProxy(config *Config) Proxy {
	mux := http.NewServeMux()
	server := &http.Server{Addr: config.Target(), Handler: mux}

	return &HttpServerProxy{Server: server, Multiplexer: mux, Backend: NewBackend(config)}
}
