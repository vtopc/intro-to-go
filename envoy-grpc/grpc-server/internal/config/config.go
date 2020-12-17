package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	LogLevel string `split_words:"true" default:"INFO"`
	Address  string `required:"true" default:":50051"`
	// DatabaseURL string `envconfig:"DATABASE_URL"`
}

// Load init config
func Load() (Config, error) {
	// TODO: add sync.Once

	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		return cfg, fmt.Errorf("failed to read config: %w", err)
	}

	return cfg, nil
}
