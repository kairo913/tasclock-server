package config

import (
	"context"
	"fmt"
	"time"

	"github.com/sethvargo/go-envconfig"
)

type SessionConfig struct {
	AccessTokenExpire time.Duration `env:"ACCESS_TOKEN_EXPIRE, default=30m"`
	RefreshTokenExpire time.Duration `env:"REFRESH_TOKEN_EXPIRE, default=720h"`
}

func NewSessionConfig(ctx context.Context) *SessionConfig {
	cfg := SessionConfig{}
	if err := envconfig.Process(ctx, &cfg); err != nil {
		panic(fmt.Sprintf("%+v", err))
	}
	return &cfg
}