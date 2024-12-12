package config

import (
	"context"
	"fmt"

	"github.com/sethvargo/go-envconfig"
)

type SQLConfig struct {
	DBType   string `env:"DB_TYPE, required"`
	Username string `env:"DB_USER, required"`
	Password string `env:"DB_PASS, required"`
	Host     string `env:"DB_HOST, required"`
	Port     string `env:"DB_PORT, required"`
	Database string `env:"DB_NAME, required"`
}

func NewSQLConfig(ctx context.Context) *SQLConfig {
	cfg := SQLConfig{}
	if err := envconfig.Process(ctx, &cfg); err != nil {
		panic(fmt.Sprintf("%+v", err))
	}
	return &cfg
}
