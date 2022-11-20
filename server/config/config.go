package config

import (
	"os"
	"strings"
)

const (
	ApiRoot  = "API_ROOT"
	Endpoint = "ENDPOINT"
	Port     = "PORT"
	FilesDir = "FILES"
)

type Config struct {
	apiRoot  string
	endpoint string
	port     string
	filesDir string
}

func NewConfig() *Config {
	return &Config{
		apiRoot:  os.Getenv(ApiRoot),
		endpoint: os.Getenv(Endpoint),
		port:     os.Getenv(Port),
		filesDir: os.Getenv(FilesDir),
	}
}

func (c *Config) GetApiPath() string {
	return strings.Join([]string{c.apiRoot, c.endpoint}, "/")
}

func (c *Config) GetPort() string {
	return c.port
}

func (c *Config) GetFilesDir() string {
	return c.filesDir
}
