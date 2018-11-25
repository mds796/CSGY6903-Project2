package proxy

import (
	"github.com/mds796/CSGY6903-Project2/crypto"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type HttpServer struct {
	DestinationTarget *url.URL
	ReverseProxy      *httputil.ReverseProxy
	UploadProxy       *httputil.ReverseProxy
	DownloadProxy     *httputil.ReverseProxy

	Cipher crypto.SymmetricCipher
}

func (s *HttpServer) Proxy() http.Handler {
	return s.ReverseProxy
}

func (s *HttpServer) Upload() http.Handler {
	return s.UploadProxy
}

func (s *HttpServer) Download() http.Handler {
	return s.DownloadProxy
}

func (s *HttpServer) WebSocket(w http.ResponseWriter, r *http.Request) {
	hijack(s.DestinationTarget, w, r)
}

func NewBackend(config *Config) Backend {
	data, err := ioutil.ReadFile(config.SymmetricKeyPath)
	if err != nil {
		panic(err)
	}

	symmetric := crypto.NewSymmetric(data)
	proxy := httputil.NewSingleHostReverseProxy(config.DestinationTarget())
	uploadProxy := &httputil.ReverseProxy{Director: encryptFiles(proxy.Director, symmetric)}
	downloadProxy := &httputil.ReverseProxy{Director: proxy.Director, ModifyResponse: decryptFiles(symmetric)}

	return &HttpServer{
		DestinationTarget: config.DestinationTarget(),
		ReverseProxy:      proxy,
		UploadProxy:       uploadProxy,
		DownloadProxy:     downloadProxy}
}
