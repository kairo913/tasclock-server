package config

import (
	"context"
	"fmt"

	"github.com/sethvargo/go-envconfig"
)

type HashConfig struct {
	SecretSalt string `env:"SECRET_SALT, required"`
	HashCount  int    `env:"HASH_COUNT, default=100000"`
}

func NewHashConfig(ctx context.Context) *HashConfig {
	cfg := HashConfig{}
	if err := envconfig.Process(ctx, &cfg); err != nil {
		panic(fmt.Sprintf("%+v", err))
	}
	return &cfg
}