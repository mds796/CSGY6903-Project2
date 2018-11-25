package proxy

import "net/http"

type HttpServer struct {
	Destination       *http.Client
	DestinationTarget string
	UploadApi         string
	DownloadApi       string
}

func (s *HttpServer) Proxy(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (s *HttpServer) Upload(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (s *HttpServer) Download(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func NewBackend(config *Config) Backend {
	client := &http.Client{}

	return &HttpServer{
		Destination:       client,
		DestinationTarget: config.DestinationTarget(),
		UploadApi:         config.UploadApi,
		DownloadApi:       config.DownloadApi}
}
