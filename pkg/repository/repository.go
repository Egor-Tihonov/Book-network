// Package repository ...
package repository

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

// PostgresDB db
type DBPostgres struct {
	Pool *pgxpool.Pool
}

// New create new connection with db
func New(connString string) (*DBPostgres, error) {
	poolP, err := pgxpool.Connect(context.Background(), connString)
	if err != nil {
		return nil, err
	}
	return &DBPostgres{Pool: poolP}, nil
}
