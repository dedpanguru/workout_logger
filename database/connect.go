package database

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

func NewConn(dbURL string) (*pgxpool.Pool, error) {
	db, err := pgxpool.Connect(context.Background(), dbURL)
	if err != nil {
		return nil, err
	}
	return db, nil
}
