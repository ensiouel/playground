package postgresql

import (
	"context"
	"fmt"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
}

func NewClient(ctx context.Context, config *Config) (pool *pgxpool.Pool, err error) {
	var (
		parseConfig *pgxpool.Config
	)

	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", config.Username, config.Password, config.Host,
		config.Port, config.Database)

	parseConfig, err = pgxpool.ParseConfig(connString)
	if err != nil {
		return
	}

	pool, err = pgxpool.ConnectConfig(ctx, parseConfig)
	if err != nil {
		return
	}

	err = pool.Ping(ctx)
	if err != nil {
		return
	}

	return
}
