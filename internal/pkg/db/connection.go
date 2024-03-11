package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lkzcover/pinread/internal/pkg/env"
)

func NewDBConn(ctx context.Context) (*pgxpool.Pool, error) {
	dbCfg, err := pgxpool.ParseConfig(env.Setup.DBConn)
	if err != nil {
		return nil, err
	}

	conn, err := pgxpool.NewWithConfig(ctx, dbCfg)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
