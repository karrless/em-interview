// Tools for work with configuration
package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/karrless/em-interview/pkg/db/postgres"
)

type Config struct {
	postgres.PostgresConfig
}

func New() *Config {
	cfg := Config{}
	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		return nil
	}
	return &cfg
}
