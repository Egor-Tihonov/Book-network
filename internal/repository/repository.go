package repository

import (
	"context"
	"fmt"

	"github.com/Egor-Tihonov/Book-network/internal/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
)

//PostgresDB db
type PostgresDB struct {
	Pool *pgxpool.Pool
}

//Create new connection with db
func New(connString string) (*PostgresDB, error) {
	poolP, err := pgxpool.Connect(context.Background(), connString)
	if err != nil {
		return nil, err
	}
	return &PostgresDB{Pool: poolP}, nil
}

//CreateUser add new user into db
func (r *PostgresDB) CreateUser(ctx context.Context, person *model.UserModel) error {
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
		return fmt.Errorf("user with this id doesnt exist")
	}
	if err != nil {
		if err == pgx.ErrNoRows {
			return fmt.Errorf("user with this id doesnt exist: %v", err)
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
		return fmt.Errorf("user with this id doesnt exist")
	}
	if err != nil {
		log.Errorf("error with update user %v", err)
		return err
	}
	return nil
}

// SelectByID : select one user by his ID
func (r *PostgresDB) SelectUser(ctx context.Context, id string) (*model.UserModel, error) {
	p := model.UserModel{}
	err := r.Pool.QueryRow(ctx, "select username,name,password from persons where id=$1", id).Scan(
		&p.Username, &p.Name, &p.Password)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("user with this id doesnt exist: %v", err)
		}
		log.Errorf("database error, select by id: %v", err)
		return nil, err /*p, fmt.errorf("user with this id doesnt exist")*/
	}
	return &p, nil
}

//SelectUserAuth
func (r *PostgresDB) SelectUserAuth(ctx context.Context, username string) (*model.UserModel, error) {
	p := model.UserModel{}
	err := r.Pool.QueryRow(ctx, "select id,username,name,password from persons where username=$1", username).Scan(
		&p.Id, &p.Username, &p.Name, &p.Password)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("user with this id doesnt exist: %v", err)
		}
		log.Errorf("database error, select by id: %v", err)
		return nil, err /*p, fmt.errorf("user with this id doesnt exist")*/
	}
	return &p, nil
}
