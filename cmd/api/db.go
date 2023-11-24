package main

import (
	"context"
	"fmt"

	"github.com/ik5/echo_api_test/utils/runtimeutils"
	"github.com/jackc/pgx/v5/pgxpool"
)

// InitPGPool initialize a pool of postgres db
func InitPGPool(conf *settings) (*pgxpool.Pool, *pgxpool.Config, error) {
	funcName := runtimeutils.GetCallerFunctionName()
	dsn := conf.PGDSN()

	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, config, fmt.Errorf("[%s] %w", funcName, err)
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		err = fmt.Errorf("[%s] %w", funcName, err)
	}

	return pool, config, err
}
