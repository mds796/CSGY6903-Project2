package proxy

import "strconv"

// Config is a configuration struct for the proxy server.
type Config struct {
	Host string
	Port uint16

	DestinationHost string
	DestinationPort uint16

	CertificatePath string
	KeyPath         string
}

func (c *Config) Target() string {
	return c.target(c.Host, c.Port)
}

func (c *Config) DestinationTarget() string {
	return c.target(c.DestinationHost, c.DestinationPort)
}

func (c *Config) target(host string, port uint16) string {
	return host + ":" + strconv.Itoa(int(port))
}
