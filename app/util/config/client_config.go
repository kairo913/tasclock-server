package config

import (
	"context"
	"fmt"

	"github.com/sethvargo/go-envconfig"
)

type ClientConfig struct {
	Port       string `env:"CLIENT_PORT, default=3000"`
}

func NewClientConfig(ctx context.Context) *ClientConfig {
	cfg := ClientConfig{}
	if err := envconfig.Process(ctx, &cfg); err != nil {
		panic(fmt.Sprintf("%+v", err))
	}
	return &cfg
}
