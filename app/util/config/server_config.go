package config

import (
	"context"
	"fmt"

	"github.com/sethvargo/go-envconfig"
)

type ServerConfig struct {
	Port           string `env:"PORT, default=8080"`
	ProductionMode bool   `env:"PRODUCTION_MODE, default=false"`
}

func NewServerConfig(ctx context.Context) *ServerConfig {
	cfg := ServerConfig{}
	if err := envconfig.Process(ctx, &cfg); err != nil {
		panic(fmt.Sprintf("%+v", err))
	}
	return &cfg
}
