// Package repository ...
package repository

import (
	"context"

	"github.com/Egor-Tihonov/Book-network/internal/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx"
	"github.com/sirupsen/logrus"
)

// Create add new user into db
func (r *PostgresDB) Create(ctx context.Context, user *model.UserModel) error {
	newID := uuid.New().String()
	_, err := r.Pool.Exec(ctx, "insert into users(id,username,name,password,email) values($1,$2,$3,$4,$5)",
		newID, &user.Username, &user.Name, &user.Password, &user.Email)
	if err != nil {
		return model.ErrorUserAlreadyExist
	}
	return nil
}

// Delete : delete user by his ID
func (r *PostgresDB) Delete(ctx context.Context, id string) error {
	a, err := r.Pool.Exec(ctx, "delete from users where id=$1", id)
	if a.RowsAffected() == 0 {
		return model.ErrorUserDoesntExist
	}
	if err != nil {
		if err == pgx.ErrNoRows {
			return model.ErrorUserDoesntExist
		}
		logrus.Errorf("error with delete user %e", err)
		return err
	}
	return nil
}

// Update update user in db
func (r *PostgresDB) Update(ctx context.Context, id string, c *model.UserUpdate) error {
	a, err := r.Pool.Exec(ctx, "update users set city=$1, country=$2, status=$3, phone=$4, bthsday=$5, email=$6 where id=$7",
		&c.City, &c.Country, &c.Status, &c.Phone, &c.Bthsday, &c.Email, id)
	if a.RowsAffected() == 0 {
		return model.ErrorUserDoesntExist
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
	err := r.Pool.QueryRow(ctx, "select username,name,status,email from users where id=$1", id).Scan(
		&p.Username, &p.Name, &p.Status, &p.Email	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, model.ErrorUserDoesntExist
		}
		return nil, err
	}
	return &p, nil
}

// GetAuth get user from db for authentication and create jwt tokens
func (r *PostgresDB) GetAuthByUsername(ctx context.Context, authString string) (*model.AuthUserModel, error) {
	var u model.AuthUserModel
	err := r.Pool.QueryRow(ctx, "select password,id,email from users where username=$1", authString).Scan(
		&u.ID, &u.ID, &u.Email)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, model.ErrorUserDoesntExist
		}
		logrus.Errorf("database error, select by id: %e", err)
		return nil, err
	}
	return &u, nil
}

func (r *PostgresDB) GetAuthByEmail(ctx context.Context, authString string) (*model.AuthUserModel, error) {
	var u model.AuthUserModel
	err := r.Pool.QueryRow(ctx, "select password,id,email from users where email=$1", authString).Scan(
		&u.Password, &u.ID, &u.Email)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, model.ErrorUserDoesntExist
		}
		logrus.Errorf("database error, select by id: %e", err)
		return nil, err
	}
	return &u, nil
}
