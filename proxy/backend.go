package proxy

import (
	"github.com/mds796/CSGY6903-Project2/crypto"
	"net/http"
	"net/http/httputil"
)

type HttpServer struct {
	ReverseProxy  *httputil.ReverseProxy
	UploadProxy   *httputil.ReverseProxy
	DownloadProxy *httputil.ReverseProxy

	Cipher crypto.SymmetricCipher
}

func (s *HttpServer) Proxy(w http.ResponseWriter, r *http.Request) {
	s.ReverseProxy.ServeHTTP(w, r)
}

func (s *HttpServer) Upload(w http.ResponseWriter, r *http.Request) {
	s.UploadProxy.ServeHTTP(w, r)
}

func (s *HttpServer) Download(w http.ResponseWriter, r *http.Request) {
	s.DownloadProxy.ServeHTTP(w, r)
}

func NewBackend(config *Config) Backend {
	symmetric := crypto.NewSymmetric("foobar")
	proxy := httputil.NewSingleHostReverseProxy(config.DestinationTarget())
	uploadProxy := &httputil.ReverseProxy{Director: encryptFiles(proxy.Director, symmetric)}
	downloadProxy := &httputil.ReverseProxy{Director: proxy.Director, ModifyResponse: decryptFiles(symmetric)}

	return &HttpServer{
		ReverseProxy:  proxy,
		UploadProxy:   uploadProxy,
		DownloadProxy: downloadProxy}
}
