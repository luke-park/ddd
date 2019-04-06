package config

import (
	"github.com/BurntSushi/toml"
)

// Config is the set of server configuration values.
type Config struct {
	Port         uint16
	SymmetricKey string
}

// Load into a Config instance.
func Load() (*Config, error) {
	config := &Config{}

	_, err := toml.DecodeFile("./ddds.toml", config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
