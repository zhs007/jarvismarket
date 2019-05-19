package jarvismarket

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// Config - config
type Config struct {
	GRPCAddr     string
	HTTPAddr     string
	Repositories []string
}

// LoadConfig - load config
func LoadConfig(filename string) (*Config, error) {
	fi, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	if err != nil {
		return nil, err
	}

	cfg := &Config{}

	err = yaml.Unmarshal(fd, cfg)
	if err != nil {
		return nil, err
	}

	err = checkConfig(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

// checkConfig -
func checkConfig(cfg *Config) error {
	if cfg.GRPCAddr == "" {
		return ErrNoServerAddress
	}

	if cfg.HTTPAddr == "" {
		return ErrNoHTTPServerAddr
	}

	if len(cfg.Repositories) <= 0 {
		return ErrNoRepositories
	}

	return nil
}
