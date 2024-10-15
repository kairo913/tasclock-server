package infra

import (
	"database/sql"

	"github.com/caarlos0/env"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"

	"github.com/kairo913/tasclock-server/app/interface/database"
)

type SqlHandler struct {
	client *sql.DB
}

type SQLConfig struct {
	DBType   string `env:"DB_TYPE" envDefault:"mysql"`
	Username string `env:"DB_USER" env:"root"`
	Password string `env:"DB_PASS" env:"password"`
	Host     string `env:"DB_HOST" envDefault:"localhost"`
	Port     string `env:"DB_PORT" envDefault:"3306"`
	Database string `env:"DB_NAME" envDefault:"tasclock"`
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

func (c *SqlHandler) Execute(statement string, args ...interface{}) (database.Result, error) {
	res := SqlResult{}
	result, err := c.client.Exec(statement, args...)
	if err != nil {
		return res, err
	}
	res.Result = result
	return res, nil
}

func (c *SqlHandler) Query(statement string, args ...interface{}) (database.Row, error) {
	rows, err := c.client.Query(statement, args...)
	if err != nil {
		return nil, err
	}
	return &SqlRow{Rows: rows}, nil
}

type SqlResult struct {
	Result sql.Result
}

func (r SqlResult) LastInsertId() (int64, error) {
	return r.Result.LastInsertId()
}

func (r SqlResult) RowsAffected() (int64, error) {
	return r.Result.RowsAffected()
}

type SqlRow struct {
	Rows *sql.Rows
}

func (r SqlRow) Scan(dest ...interface{}) error {
	return r.Rows.Scan(dest...)
}

func (r SqlRow) Next() bool {
	return r.Rows.Next()
}

func (r SqlRow) Close() error {
	return r.Rows.Close()
}
