package infra

import (
	"context"
	"time"

	"github.com/caarlos0/env"
	"github.com/redis/go-redis/v9"
)

type RedisHandler struct {
	client *redis.Client
}

type RedisConfig struct {
	Addr     string `env:"REDIS_ADDR" envDefault:"redis:6379"`
	PassWord string `env:"REDIS_PASSWORD" envDefault:""`
	DB       int    `env:"REDIS_DB" envDefault:"0"`
}

func NewRedisHandler() (*RedisHandler, error) {
	cfg := RedisConfig{}
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.PassWord,
		DB:       cfg.DB,
	})

	return &RedisHandler{client: client}, nil
}

func (r *RedisHandler) Close() error {
	return r.client.Close()
}

func (r *RedisHandler) Set(key, value string, exp time.Duration) error {
	return r.client.Set(context.Background(), key, value, exp).Err()
}

func (r *RedisHandler) Get(key string) (string, error) {
	return r.client.Get(context.Background(), key).Result()
}

func (r *RedisHandler) Del(key string) error {
	return r.client.Del(context.Background(), key).Err()
}

func (r *RedisHandler) Exists(key string) (bool, error) {
	res, err := r.client.Exists(context.Background(), key).Result()
	if err != nil {
		return false, err
	}

	if res == 1 {
		return true, nil
	}

	return false, nil
}

func (r *RedisHandler) FlushAll() error {
	return r.client.FlushAll(context.Background()).Err()
}