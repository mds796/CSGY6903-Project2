package proxy

import (
	"log"
	"net/http"
)

type HttpServerProxy struct {
	Multiplexer *http.ServeMux
	Server      *http.Server

	Backend Backend

	CertificatePath string
	KeyPath         string

	UploadApi    string
	DownloadApi  string
	WebSocketApi string
}

func (p *HttpServerProxy) Start() error {
	p.configureRoutes()
	return p.listenAndServe()
}

func (p *HttpServerProxy) configureRoutes() {
	p.Multiplexer.Handle(p.UploadApi, p.Backend.Upload())
	p.Multiplexer.Handle(p.DownloadApi, p.Backend.Download())
	p.Multiplexer.HandleFunc(p.WebSocketApi, p.Backend.WebSocket)
	p.Multiplexer.Handle("/", p.Backend.Proxy())
}

func (p *HttpServerProxy) listenAndServe() error {
	log.Printf("Proxy server now listening on %v.\n", p.Server.Addr)

	return p.Server.ListenAndServeTLS(p.CertificatePath, p.KeyPath)
}

func (p *HttpServerProxy) Stop() error {
	return p.Server.Close()
}

func NewProxy(config *Config) Proxy {
	mux := http.NewServeMux()
	server := &http.Server{Addr: config.Target(), Handler: mux}

	return &HttpServerProxy{
		Server:          server,
		Multiplexer:     mux,
		Backend:         NewBackend(config),
		CertificatePath: config.CertificatePath,
		KeyPath:         config.KeyPath,
		UploadApi:       config.UploadApi,
		DownloadApi:     config.DownloadApi,
		WebSocketApi:    config.WebSocketApi}
}
