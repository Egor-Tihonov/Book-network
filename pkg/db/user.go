// Package repository ...
package db

import (
	"context"
	"fmt"
	"time"

	"github.com/Egor-Tihonov/Book-network/pkg/models"
	"github.com/jackc/pgx"
	"github.com/sirupsen/logrus"
)

// Create add new user into db
func (p *DBPostgres) Create(ctx context.Context, user *models.UserModel) error {
	date := time.Now()
	_, err := p.Pool.Exec(ctx, "insert into users(id,username,name,email,joinDate) values($1,$2,$3,$4,$5)",
		&user.ID, &user.Username, &user.Name, &user.Email, &date)
	if err != nil {
		return models.ErrorUserAlreadyExist
	}
	return nil
}

func (p *DBPostgres) GetUser(ctx context.Context, id string) (*models.User, error) {
	user := models.User{}
	date := time.Time{}
	err := p.Pool.QueryRow(ctx, "select username,name,status,email,joinDate from users where id=$1", id).Scan(
		&user.Username, &user.Name, &user.Status, &user.Email, &date)
	if err != nil {
		if err.Error() == pgx.ErrNoRows.Error() {
			return nil, models.ErrorUserDoesntExist
		}
		return nil, err
	}
	user.JoinDate = date.Format("2006-01-02")
	return &user, nil
}

// Delete : delete user by his ID
func (p *DBPostgres) Delete(ctx context.Context, id string) error {
	a, err := p.Pool.Exec(ctx, "delete from users where id=$1", id)
	if err != nil {
		if err.Error() == pgx.ErrNoRows.Error() {
			return models.ErrorUserDoesntExist
		}
		logrus.Errorf("error with delete user %w", err)
		return err
	}
	if a.RowsAffected() == 0 {
		return models.ErrorUserDoesntExist
	}
	return nil
}

// func (p *DBPostgres) GetUserForUpdate(ctx context.Context, id string) (*models.UserUpdate, error) {
// 	user := models.UserUpdate{}
// 	err := p.Pool.QueryRow(ctx, "select status,name,username,password from users where id = $1", id).Scan(
// 		&user.Status, &user.Name, &user.Username, &user.Password)
// 	if err != nil {
// 		if err.Error() == pgx.ErrNoRows.Error() {
// 			return nil, models.ErrorUserDoesntExist
// 		} else {
// 			return nil, err
// 		}
// 	}

// 	return &user, err
// }

func (p *DBPostgres) GetLastUsersIDs(ctx context.Context) ([]*models.LastUsers, error) {
	var lastUsers []*models.LastUsers
	sql := "select id,username" +
		" from users where EXTRACT(YEAR FROM joindate) = EXTRACT(YEAR FROM NOW()) AND EXTRACT(WEEK FROM joindate) = EXTRACT (WEEK FROM NOW()) order by joindate desc "
	rows, err := p.Pool.Query(ctx, sql)
	if err != nil {
		if err.Error() == pgx.ErrNoRows.Error() {
			return nil, models.ErrorNoPosts
		}
		logrus.Errorf("database error with select all users id, %w", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		lastUser := models.LastUsers{}
		err = rows.Scan(&lastUser.Id, &lastUser.Username)
		if err != nil {
			logrus.Errorf("database error with select all users id, %w", err)
			return nil, err
		}
		lastUsers = append(lastUsers, &lastUser)
	}
	return lastUsers, err
}

// Update update user in db
func (p *DBPostgres) Update(ctx context.Context, id string, c *models.UserUpdate) error {
	fmt.Println(id)
	a, err := p.Pool.Exec(ctx, "update users set status=$1, name=$2, username=$3 where id=$4",
		&c.Status, &c.Name, &c.Username, id)

	if err != nil {
		logrus.Errorf("error with update user %w", err)
		return err
	}

	if a.RowsAffected() == 0 {
		logrus.Errorf("error with update user %w", err)
		return models.ErrorUserDoesntExist
	}
	return nil
}

func (p *DBPostgres) AddSubscriprion(ctx context.Context, subid, id string) error {
	a, err := p.Pool.Exec(ctx, "update users set subscriptions = ARRAY_APPEND(subscriptions,$1) where id = $2", subid, id)
	if a.RowsAffected() == 0 {
		return fmt.Errorf("users doesnt exist, failing add subscription")
	}
	if err != nil {
		logrus.Errorf("error while add subscription, %w", err)
		return err
	}
	return nil
}

func (p *DBPostgres) DeleteSubscription(ctx context.Context, subid, id string) error {
	a, err := p.Pool.Exec(ctx, "update users set subscriptions = array_remove(subscriptions,$1) where id = $2", subid, id)
	if a.RowsAffected() == 0 {
		return fmt.Errorf("users doesnt exist, failing remove subscription")
	}
	if err != nil {
		logrus.Errorf("error while remove subscription, %w", err)
		return err
	}
	return nil
}

func (p *DBPostgres) CheckSubs(ctx context.Context, id, userid string) (bool, error) {
	var position int
	err := p.Pool.QueryRow(ctx, "select array_position(subscriptions, $1)::INTEGER from users where id=$2", id, userid).Scan(
		&position)
	if err != nil {
		if position == 0 {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func (p *DBPostgres) GetMySubs(ctx context.Context, id string) ([]*models.User, error) {
	var subusers []*models.User
	date := time.Time{}
	rows, err := p.Pool.Query(ctx, "select id,name,username,status,email,joindate from users where id IN "+
		"(select unnest(subscriptions) from users where id=$1)", id)
	if err != nil {
		if err.Error() == pgx.ErrNoRows.Error() {
			return nil, models.ErrorNoPosts
		}
		logrus.Errorf("database error with select sub users id, %w", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		user := models.User{}
		err = rows.Scan(&user.ID, &user.Name, &user.Username, &user.Status, &user.Email, &date)
		if err != nil {
			logrus.Errorf("database error with select sub users id, %w", err)
			return nil, err
		}
		user.JoinDate = date.Format("2006-01-02")
		subusers = append(subusers, &user)
	}
	return subusers, err

}
