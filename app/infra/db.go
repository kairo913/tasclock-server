package infra

import (
	"database/sql"

	"github.com/caarlos0/env"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

type SqlHandler struct {
	client *sql.DB
}

type SQLConfig struct {
	DBType   string `env:"DB_TYPE"`
	Username string `env:"DB_USER"`
	Password string `env:"DB_PASS"`
	Host     string `env:"DB_HOST"`
	Port     string `env:"DB_PORT"`
	Database string `env:"DB_NAME"`
}

func NewSqlHandler() (*SqlHandler, error) {
	cfg := SQLConfig{}
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	client, err := sql.Open(cfg.DBType, cfg.Username+":"+cfg.Password+"@tcp("+cfg.Host+":"+cfg.Port+")/"+cfg.Database)
	if err != nil {
		return nil, err
	}

	return &SqlHandler{client: client}, nil
}

func (c *SqlHandler) Close() error {
	return c.client.Close()
}
