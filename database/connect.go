package database

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

func NewConnPool(dbURL string) (*pgxpool.Pool, context.Context, error) {
	ctx := context.Background()
	dbpool, err := pgxpool.Connect(ctx, dbURL)
	if err != nil {
		return nil, nil, err
	}
	return dbpool, ctx, nil
}
