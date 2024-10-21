package database

import (
	"context"
	"database/sql"

	"github.com/kairo913/tasclock-server/app/infra/repository"
	"github.com/kairo913/tasclock-server/app/util/config"
)

type SqlHandler struct {
	Client *sql.DB
}

func NewSqlHandler(ctx context.Context) (*SqlHandler, error) {
	cfg := config.NewSQLConfig(ctx)

	client, err := sql.Open(cfg.DBType, cfg.Username+":"+cfg.Password+"@tcp("+cfg.Host+":"+cfg.Port+")/"+cfg.Database)
	if err != nil {
		return nil, err
	}

	return &SqlHandler{Client: client}, nil
}

func (c *SqlHandler) Close() error {
	return c.Client.Close()
}

func (c *SqlHandler) Execute(statement string, args ...interface{}) (repository.Result, error) {
	res := SqlResult{}
	result, err := c.Client.Exec(statement, args...)
	if err != nil {
		return res, err
	}
	res.Result = result
	return res, nil
}

func (c *SqlHandler) Query(statement string, args ...interface{}) (repository.Row, error) {
	rows, err := c.Client.Query(statement, args...)
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
