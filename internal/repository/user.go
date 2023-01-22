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
func (p *PostgresDB) Create(ctx context.Context, user *model.UserModel,date time.Time) error {
	newID := uuid.New().String()
	_, err := p.Pool.Exec(ctx, "insert into users(id,username,name,password,email,joinDate) values($1,$2,$3,$4,$5,$6)",
		newID, &user.Username, &user.Name, &user.Password, &user.Email, &date)
	if err != nil {
		return model.ErrorUserAlreadyExist
	}
	return nil
}

// Delete : delete user by his ID
func (p *PostgresDB) Delete(ctx context.Context, id string) error {
	a, err := p.Pool.Exec(ctx, "delete from users where id=$1", id)
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

func (p *PostgresDB) GetUserForUpdate(ctx context.Context, id string) (*model.UserUpdate, error) {
	user := model.UserUpdate{}
	err := p.Pool.QueryRow(ctx, "select status,name,username,password from users where id = $1", id).Scan(
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
func (p *PostgresDB) Update(ctx context.Context, id string, c *model.UserUpdate) error {
	a, err := p.Pool.Exec(ctx, "update users set status=$1, name=$2, username=$3,password=$4 where id=$5",
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
func (p *PostgresDB) Get(ctx context.Context, id string) (*model.User, error) {
	user := model.User{}
	date := time.Time{}
	err := p.Pool.QueryRow(ctx, "select username,name,status,email,joinDate from users where id=$1", id).Scan(
		&user.Username, &user.Name, &user.Status, &user.Email, &date)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, model.ErrorUserDoesntExist
		}
		return nil, err
	}
	user.JoinDate = date.Format("2006-01-02")
	return &user, nil
}

// GetAuth get user from db for authentication and create jwt tokens
func (p *PostgresDB) GetAuthByUsername(ctx context.Context, authString string) (*model.AuthUserModel, error) {
	var u model.AuthUserModel
	err := p.Pool.QueryRow(ctx, "select password,id,email from users where username=$1", authString).Scan(
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

func (p *PostgresDB) GetAuthByEmail(ctx context.Context, authString string) (*model.AuthUserModel, error) {
	var u model.AuthUserModel
	err := p.Pool.QueryRow(ctx, "select password,id,email from users where email=$1", authString).Scan(
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
