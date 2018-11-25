package proxy

import (
	"net/url"
	"strconv"
)

// Config is a configuration struct for the proxy server.
type Config struct {
	Host string
	Port uint16

	DestinationScheme string
	DestinationHost   string
	DestinationPort   uint16

	UploadApi    string
	DownloadApi  string
	WebSocketApi string

	CertificatePath  string
	KeyPath          string
	SymmetricKeyPath string
}

func (c *Config) Target() string {
	return c.target(c.Host, c.Port)
}

func (c *Config) DestinationTarget() *url.URL {
	return &url.URL{Scheme: c.DestinationScheme, Host: c.target(c.DestinationHost, c.DestinationPort)}
}

func (c *Config) target(host string, port uint16) string {
	return host + ":" + strconv.Itoa(int(port))
}
