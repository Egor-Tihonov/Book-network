// Package repository ...
package repository

import (
	"context"
	"time"

	"github.com/Egor-Tihonov/Book-network/internal/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx"
	"github.com/sirupsen/logrus"
)

// Create add new user into db
func (r *PostgresDB) Create(ctx context.Context, user *model.UserModel) error {
	newID := uuid.New().String()
	date := time.Now()
	_, err := r.Pool.Exec(ctx, "insert into users(id,username,name,password,email,joinDate) values($1,$2,$3,$4,$5,$6)",
		newID, &user.Username, &user.Name, &user.Password, &user.Email, &date)
	if err != nil {
		return model.ErrorUserAlreadyExist
	}
	return nil
}

// Delete : delete user by his ID
func (r *PostgresDB) Delete(ctx context.Context, id string) error {
	a, err := r.Pool.Exec(ctx, "delete from users where id=$1", id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return model.ErrorUserDoesntExist
		}
		logrus.Errorf("error with delete user %e", err)
		return err
	}
	if a.RowsAffected() == 0 {
		return model.ErrorUserDoesntExist
	}	
	return nil
}

func (r *PostgresDB) GetUserForUpdate(ctx context.Context, id string) (*model.UserUpdate, error) {
	user := model.UserUpdate{}
	err := r.Pool.QueryRow(ctx, "select status,name,username,password from users where id = $1", id).Scan(
		&user.Status, &user.Name, &user.Username, &user.Password)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, model.ErrorUserDoesntExist
		} else {
			return nil, err
		}
	}

	return &user, err
}

// Update update user in db
func (r *PostgresDB) Update(ctx context.Context, id string, c *model.UserUpdate) error {
	a, err := r.Pool.Exec(ctx, "update users set status=$1, name=$2, username=$3,password=$4 where id=$5",
		&c.Status, &c.Name, &c.Username, &c.Password, id)
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
	date := time.Time{}
	err := r.Pool.QueryRow(ctx, "select username,name,status,email,joinDate from users where id=$1", id).Scan(
		&p.Username, &p.Name, &p.Status, &p.Email, &date)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, model.ErrorUserDoesntExist
		}
		return nil, err
	}
	p.JoinDate = date.Format("2006-01-02")
	return &p, nil
}

// GetAuth get user from db for authentication and create jwt tokens
func (r *PostgresDB) GetAuthByUsername(ctx context.Context, authString string) (*model.AuthUserModel, error) {
	var u model.AuthUserModel
	err := r.Pool.QueryRow(ctx, "select password,id,email from users where username=$1", authString).Scan(
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
