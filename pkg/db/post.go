package db

import (
	"context"
	"errors"
	"time"

	"github.com/Egor-Tihonov/Book-network/pkg/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
)

func (p *DBPostgres) DeletePost(ctx context.Context, postid, userid string) error {
	_, err := p.Pool.Exec(ctx, "delete * from posts where postid = $1 and userid= $2", postid, userid)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.ErrorUserDoesntExist
		}
		logrus.Errorf("user service error: delete post %w", err)
		return err
	}

	return nil
}

func (p *DBPostgres) GetMyPosts(ctx context.Context, userid string) ([]*models.Post, error) {
	var posts []*models.Post
	sql := "select posts.content, posts.id,posts.book_title,posts.author_name,posts.author_surname from posts where posts.userid=$1 order by dt_create desc"
	rows, err := p.Pool.Query(ctx, sql, userid)
	if err != nil {
		if err.Error() == pgx.ErrNoRows.Error() {
			return nil, models.ErrorNoPosts
		}
		logrus.Errorf("database error with select all posts, %w", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		pb := models.Post{}
		err = rows.Scan(&pb.Content, &pb.PostId, &pb.Title, &pb.AuthorName, &pb.AuthorSurname)
		if err != nil {
			logrus.Errorf("database error with select all posts, %w", err)
			return nil, err
		}
		posts = append(posts, &pb)
	}
	return posts, err
}

func (p *DBPostgres) CreatePost(ctx context.Context, userid, bookId string, post *models.Post, date time.Time) error {
	postId := uuid.New().String()
	sql := "insert into posts(userid, idbook, content, id, dt_create,book_title,author_name,author_surname) values($1, $2, $3, $4, $5,$6,$7,$8)"
	_, err := p.Pool.Exec(ctx, sql, userid, bookId, post.Content, postId, date, post.Title, post.AuthorName, post.AuthorSurname)
	if err != nil {
		return err
	}
	return nil
}

func (p *DBPostgres) GetForCheckPosts(ctx context.Context, userId string) ([]string, error) {
	var ids []string
	sql := "select idbook from posts where userid = $1"
	rows, err := p.Pool.Query(ctx, sql, userId)
	if err != nil {
		if err.Error() == pgx.ErrNoRows.Error() {
			return nil, nil
		}
		logrus.Errorf("database error with select all posts, %w", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		i := ""
		err = rows.Scan(&i)
		if err != nil {
			logrus.Errorf("database error with select all posts, %w", err)
			return nil, err
		}
		ids = append(ids, i)
	}
	return ids, err
}

func (p *DBPostgres) GetPost(ctx context.Context, postid string) (*models.Post, error) {
	post := models.Post{}
	err := p.Pool.QueryRow(ctx, "select posts.author_name,posts.author_surname,posts.content,posts.id,posts.book_title from posts "+
		" where posts.id=$1 order by dt_create desc", postid).Scan(
		&post.AuthorName, &post.AuthorSurname, &post.Content, &post.PostId, &post.Title)
	if err != nil {
		if err.Error() == pgx.ErrNoRows.Error() {
			return nil, models.ErrorNoPosts
		}
		logrus.Errorf("user service error: get post, %w", err)
		return nil, err
	}
	return &post, err
}

func (p *DBPostgres) GetPosts(ctx context.Context, id string) ([]*models.Feed, error) {
	var posts []*models.Feed

	sql := "select users.id, users.username,users.status, posts.dt_create,posts.content,posts.book_title,posts.author_name,posts.author_surname from users inner join posts on " +
		" users.id = posts.userid where posts.userid " +
		"= $1 or posts.userid in " +
		"(select unnest(subscriptions) from users where id=$1) " +
		"and users.id in (select unnest(subscriptions) from users where id=$1) " +
		"order by dt_create desc"
	rows, err := p.Pool.Query(ctx, sql, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, models.ErrorNoPosts
		}
		logrus.Errorf("database error with select all posts, %e", err.Error())
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		pb := models.Feed{}
		date := time.Time{}
		err = rows.Scan(&pb.UserId, &pb.Username, &pb.Status, &date, &pb.Content, &pb.Title, &pb.AuthorName, &pb.AuthorSurname)
		if err != nil {
			logrus.Errorf("database error with select all posts, %w", err)
			return nil, err
		}

		diff := time.Now().Sub(date)
		out := time.Time{}.Add(diff)

		pb.CreateDate = out.Format("15:04:05")

		posts = append(posts, &pb)
	}
	return posts, err
}
