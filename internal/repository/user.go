// Package repository ...
package repository

import (
	"context"

	"github.com/Egor-Tihonov/Book-network/internal/model"
	"github.com/Egor-Tihonov/Book-network/internal/errmodel"
	"github.com/google/uuid"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

// PostgresDB db
type PostgresDB struct {
	Pool *pgxpool.Pool
}

var (
//ErrorUserDoesntExist ...

)

// New create new connection with db
func New( /*connString string*/ ) (*PostgresDB, error) {
	poolP, err := pgxpool.Connect(context.Background(), "postgresql://postgres:123@localhost:5432/booknetwork") //connString)
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
		return user_errors.ErrorUserAlreadyExist
	}
	return nil
}

// Delete : delete user by his ID
func (r *PostgresDB) Delete(ctx context.Context, id string) error {
	a, err := r.Pool.Exec(ctx, "delete from persons where id=$1", id)
	if a.RowsAffected() == 0 {
		return user_errors.ErrorUserDoesntExist
	}
	if err != nil {
		if err == pgx.ErrNoRows {
			return user_errors.ErrorUserDoesntExist
		}
		logrus.Errorf("error with delete user %e", err)
		return err
	}
	return nil
}

// Update update user in db
func (r *PostgresDB) Update(ctx context.Context, id string, c *model.UserUpdate) error {
	a, err := r.Pool.Exec(ctx, "update persons set city=$1, country=$2, status=$3, phone=$4, bthsday=$5, email=$6 where id=$7",
		&c.City, &c.Country, &c.Status, &c.Phone, &c.Bthsday, &c.Email, id)
	if a.RowsAffected() == 0 {
		return user_errors.ErrorUserDoesntExist
	}
	if err != nil {
		logrus.Errorf("error with update user %e", err)
		return err
	}
	return nil
}

// Get : select one user by his ID
func (r *PostgresDB) Get(ctx context.Context, id string) (*model.User, error) {
	p := model.User{}
	err := r.Pool.QueryRow(ctx, "select username,name,status from persons where id=$1", id).Scan(
		&p.Username, &p.Name, &p.Status)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, user_errors.ErrorUserDoesntExist
		}
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
			return nil, user_errors.ErrorUserDoesntExist
		}
		logrus.Errorf("database error, select by id: %e", err)
		return nil, err
	}
	return &p, nil
}
