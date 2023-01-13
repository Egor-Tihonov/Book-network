// Package repository ...
package repository

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

// PostgresDB db
type PostgresDB struct {
	Pool *pgxpool.Pool
}

// New create new connection with db
func New(connString string) (*PostgresDB, error) {
	poolP, err := pgxpool.Connect(context.Background(), connString) /*"postgresql://postgres:123@localhost:5432/booknetwork"*/
	if err != nil {
		return nil, err
	}
	return &PostgresDB{Pool: poolP}, nil
}
