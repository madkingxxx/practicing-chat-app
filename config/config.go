package config

import (
	"context"
	"log"

	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	Postgres PG
}

func Load() *Config {
	cfg := &Config{}
	if err := envconfig.Process(context.Background(), cfg); err != nil {
		log.Fatal(err)
	}
	return cfg
}

type PG struct {
	URL string `env:"PG_URL,default=postgres://postgres:postgres@localhost:5432/web_socket?sslmode=disable"`
}
