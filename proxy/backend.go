package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

type HttpServer struct {
	Destination       *http.Client
	DestinationTarget *url.URL
	ReverseProxy      *httputil.ReverseProxy
	UploadApi         *url.URL
	DownloadApi       *url.URL
}

func (s *HttpServer) Proxy(w http.ResponseWriter, r *http.Request) {
	s.ReverseProxy.ServeHTTP(w, r)
}

func (s *HttpServer) WebSocket(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (s *HttpServer) Upload(w http.ResponseWriter, r *http.Request) {
	s.ReverseProxy.ServeHTTP(w, r)
}

func (s *HttpServer) Download(w http.ResponseWriter, r *http.Request) {
	s.ReverseProxy.ServeHTTP(w, r)
}

func NewBackend(config *Config) Backend {
	client := &http.Client{}

	upload, err := url.Parse(config.UploadApi)
	if err != nil {
		panic(err)
	}

	download, err := url.Parse(config.DownloadApi)
	if err != nil {
		panic(err)
	}

	return &HttpServer{
		Destination:       client,
		DestinationTarget: config.DestinationTarget(),
		UploadApi:         upload,
		DownloadApi:       download,
		ReverseProxy:      httputil.NewSingleHostReverseProxy(config.DestinationTarget())}
}
