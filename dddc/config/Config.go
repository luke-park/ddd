package config

import (
	"github.com/BurntSushi/toml"
)

// Config is the set of client configuration values.
type Config struct {
	SymmetricKey string
}

// Load into a Config instance.
func Load() (*Config, error) {
	config := &Config{}

	_, err := toml.DecodeFile("./dddc.toml", config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
