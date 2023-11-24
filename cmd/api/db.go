package main

import (
	stdcontext "context"
	"fmt"

	"github.com/ik5/echo_api_test/utils/runtimeutils"
	"github.com/jackc/pgx/v5/pgxpool"
)

// initPGPool initialize a pool of postgres db
func initPGPool(conf *settings) (*pgxpool.Pool, *pgxpool.Config, error) {
	funcName := runtimeutils.GetCallerFunctionName()
	dsn := conf.PGDSN()

	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, config, fmt.Errorf("[%s] %w", funcName, err)
	}

	pool, err := pgxpool.NewWithConfig(stdcontext.Background(), config)
	if err != nil {
		err = fmt.Errorf("[%s] %w", funcName, err)
	}

	return pool, config, err
}
