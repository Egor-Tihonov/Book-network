// Package repository ...
package repository

import (
	"context"
	"errors"

	"github.com/Egor-Tihonov/Book-network/internal/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
)

// PostgresDB db
type PostgresDB struct {
	Pool *pgxpool.Pool
}

var (
	ErrorUserDoesntExist = errors.New("user with this id/username doesnt exist")
)

// New create new connection with db
func New(connString string) (*PostgresDB, error) {
	poolP, err := pgxpool.Connect(context.Background(), connString)
	if err != nil {
		return nil, err
	}
	return &PostgresDB{Pool: poolP}, nil
}

// Create add new user into db
func (r *PostgresDB) Create(ctx context.Context, person *model.UserModel) error {
	newID := uuid.New().String()
	_, err := r.Pool.Exec(ctx, "insert into persons(id,username,name,password) values($1,$2,$3,$4)",
		newID, &person.Username, &person.Name, &person.Password)
	if err != nil {
		log.Errorf("database error with create user: %v", err)
		return err
	}
	return nil
}

// Delete : delete user by his ID
func (r *PostgresDB) Delete(ctx context.Context, id string) error {
	a, err := r.Pool.Exec(ctx, "delete from persons where id=$1", id)
	if a.RowsAffected() == 0 {
		return ErrorUserDoesntExist
	}
	if err != nil {
		if err == pgx.ErrNoRows {
			return ErrorUserDoesntExist
		}
		log.Errorf("error with delete user %v", err)
		return err
	}
	return nil
}

// Update update user in db
func (r *PostgresDB) Update(ctx context.Context, id string, p *model.UserModel) error {
	a, err := r.Pool.Exec(ctx, "update persons set username=$1,name=$2 where id=$4", &p.Username, &p.Name, id)
	if a.RowsAffected() == 0 {
		return ErrorUserDoesntExist
	}
	if err != nil {
		log.Errorf("error with update user %v", err)
		return err
	}
	return nil
}

// Get : select one user by his ID
func (r *PostgresDB) Get(ctx context.Context, id string) (*model.UserModel, error) {
	p := model.UserModel{}
	err := r.Pool.QueryRow(ctx, "select username,name from persons where id=$1", id).Scan(
		&p.Username, &p.Name)
	if err != nil {
		if err == pgx.ErrNoRows {

			return nil, ErrorUserDoesntExist
		}
		log.Errorf("database error, select by id: %v", err)
		return nil, err
	}
	return &p, nil
}

// GetAuth get user from db for authentication and create jwt tokens
func (r *PostgresDB) GetAuth(ctx context.Context, username string) (*model.UserModel, error) {
	p := model.UserModel{}
	err := r.Pool.QueryRow(ctx, "select id,username,password from persons where username=$1", username).Scan(
		&p.ID, &p.Username, &p.Password)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, ErrorUserDoesntExist
		}
		log.Errorf("database error, select by id: %v", err)
		return nil, err
	}
	return &p, nil
}
