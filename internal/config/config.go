package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/karrless/em-interview/pkg/db/postgres"
)

type Config struct {
	postgres.PostgresConfig
}

func New(path string) *Config {
	cfg := Config{}
	if path == "" {
		path = "./.env"
	}
	err := cleanenv.ReadConfig(path, &cfg)
	if err != nil {
		return nil
	}
	return &cfg
}
