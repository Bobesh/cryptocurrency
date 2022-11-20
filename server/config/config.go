package config

import (
	"os"
	"strings"
)

const (
	ApiRoot  = "API_ROOT"
	Endpoint = "ENDPOINT"
	Port     = "PORT"
)

type Config struct {
	apiRoot  string
	endpoint string
	port     string
}

func NewConfig() *Config {
	return &Config{
		apiRoot:  os.Getenv(ApiRoot),
		endpoint: os.Getenv(Endpoint),
		port:     os.Getenv(Port),
	}
}

func (c *Config) GetApiPath() string {
	return strings.Join([]string{c.apiRoot, c.endpoint}, "/")
}

func (c *Config) GetPort() string {
	return c.port
}
