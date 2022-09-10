// Package repository : file contains operations with PostgresDB
package repository

import (
	"context"
	"fmt"

	"github.com/Egor-Tihonov/Book-network/internal/model"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
)

// PRepository :creating new connection with PostgreDB
type PRepository struct {
	Pool *pgxpool.Pool
}

// Create : insert new user into database
func (r *PRepository) Create(ctx context.Context, person *model.Person) (string, error) {
	newID := uuid.New().String()
	_, err := r.Pool.Exec(ctx, "insert into persons(id,name,works,age,password) values($1,$2,$3,$4,$5)",
		newID, &person.Name, &person.Works, &person.Age, &person.Password)
	if err != nil {
		log.Errorf("database error with create user: %v", err)
		return "", err
	}
	return newID, nil
}

// Delete : delete user by his ID
func (r *PRepository) Delete(ctx context.Context, id string) error {
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

// UpdateAuth : update user refreshToken by his ID
func (r *PRepository) UpdateAuth(ctx context.Context, id, refreshToken string) error {
	a, err := r.Pool.Exec(ctx, "update persons set refreshToken=$1 where id=$2", refreshToken, id)
	if a.RowsAffected() == 0 {
		return fmt.Errorf("user with this id doesnt exist")
	}
	if err != nil {
		log.Errorf("error with update user %v", err)
		return err
	}
	return nil
}

// Update update user in db
func (r *PRepository) Update(ctx context.Context, id string, p *model.Person) error {
	a, err := r.Pool.Exec(ctx, "update persons set name=$1,works=$2,age=$3 where id=$4", &p.Name, &p.Works, &p.Age, id)
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
func (r *PRepository) SelectByID(ctx context.Context, id string) (model.Person, error) {
	p := model.Person{}
	err := r.Pool.QueryRow(ctx, "select id,name,works,age,password from persons where id=$1", id).Scan(
		&p.ID, &p.Name, &p.Works, &p.Age, &p.Password)
	if err != nil {
		if err == pgx.ErrNoRows {
			return model.Person{}, fmt.Errorf("user with this id doesnt exist: %v", err)
		}
		log.Errorf("database error, select by id: %v", err)
		return model.Person{}, err /*p, fmt.errorf("user with this id doesnt exist")*/
	}
	return p, nil
}

// SelectByIDAuth select auth user
func (r *PRepository) SelectByIDAuth(ctx context.Context, id string) (model.Person, error) {
	p := model.Person{}
	err := r.Pool.QueryRow(ctx, "select id,refreshToken from persons where id=$1", id).Scan(&p.ID, &p.RefreshToken)

	if err != nil /*err==no-records*/ {
		if err == pgx.ErrNoRows {
			return model.Person{}, fmt.Errorf("user with this id doesnt exist: %v", err)
		}
		log.Errorf("database error, select by id: %v", err)
		return model.Person{}, err /*p, fmt.errorf("user with this id doesnt exist")*/
	}
	return p, nil
}
