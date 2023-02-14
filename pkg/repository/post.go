package repository

import (
	"context"
	"errors"
	"time"

	"github.com/Egor-Tihonov/Book-network/pkg/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
)

func (p *DBPostgres) GetAll(ctx context.Context, userid string) ([]*models.Post, error) {
	var posts []*models.Post
	sql := "select author.name,author.surname,books.title,posts.content,posts.id from books inner join author on author.id=books.idauthor" +
		" inner join posts on books.id=posts.idbook where posts.userid=$1 order by dt_create desc"
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
		po := models.Post{}
		err = rows.Scan(&po.AuthorName, &po.AuthorSurname, &po.Title, &po.Content, &po.PostId)
		if err != nil {
			logrus.Errorf("database error with select all posts, %w", err)
			return nil, err
		}
		posts = append(posts, &po)
	}
	return posts, err
}

func (p *DBPostgres) CreatePost(ctx context.Context, userid, authorId, bookId string, post *models.Post, date time.Time) error {
	if authorId == "" {
		authorId = uuid.New().String()
		sql := "insert into author(id, name, surname) values($1,$2,$3)"
		_, err := p.Pool.Exec(ctx, sql, authorId, post.AuthorName, post.AuthorSurname)
		if err != nil {
			return err
		}
	}
	if bookId == "" {
		bookId = uuid.New().String()
		sql := "insert into books(id, idauthor, title) values($1,$2,$3)"
		_, err := p.Pool.Exec(ctx, sql, bookId, authorId, post.Title)
		if err != nil {
			return err
		}
	}
	postId := uuid.New().String()
	sql := "insert into posts(userid, idbook, content, id, dt_create) values($1, $2, $3, $4, $5)"
	_, err := p.Pool.Exec(ctx, sql, userid, bookId, post.Content, postId, date)
	if err != nil {
		return err
	}
	return nil
}

func (p *DBPostgres) GetForCheckAuthor(ctx context.Context, name, surname string) (string, string, error) {
	authorid := ""
	booksid := ""
	sql := "select author.id,books.id from author inner join books on books.idauthor=author.id inner join posts on posts.idbook = books.id where author.name =$1 and author.surname = $2"
	err := p.Pool.QueryRow(ctx, sql, name, surname).Scan(&authorid, &booksid)
	if err != nil {
		if err.Error() == pgx.ErrNoRows.Error() {
			return booksid, authorid, nil
		}
		logrus.Errorf("error get author from db, %w", err)
		return booksid, authorid, err
	}
	return booksid, authorid, nil
}

func (p *DBPostgres) GetForCheckPosts(ctx context.Context, userId string) ([]string, error) {
	var ids []string
	sql := "select idbook from posts where userid = $1"
	rows, err := p.Pool.Query(ctx, sql, userId)
	if err != nil {
		if err.Error() == pgx.ErrNoRows.Error() {
			return nil, models.ErrorNoPosts
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

func (p *DBPostgres) GetPost(ctx context.Context, userid, postid string) (*models.Post, error) {
	post := models.Post{}
	err := p.Pool.QueryRow(ctx, "select author.name,author.surname,posts.content,posts.id,books.title from books inner join author on author.id=books.idauthor"+
		" inner join posts on books.id=posts.idbook where posts.userid=$1 and posts.id=$2 order by dt_create desc", userid, postid).Scan(
		&post.AuthorName, &post.AuthorSurname, &post.Content, &post.PostId, &post.Title)
	if err != nil {
		if err.Error() == pgx.ErrNoRows.Error() {
			return nil, models.ErrorNoPosts
		}
		return nil, err
	}
	return &post, err
}

func (p *DBPostgres) GetLast(ctx context.Context, userid string) ([]*models.LastPost, error) {
	var lastPosts []*models.LastPost
	sql := "select posts.id,books.title from books inner join author on author.id=books.idauthor" +
		" inner join posts on books.id=posts.idbook where posts.userid=$1 and EXTRACT(YEAR FROM dt_create) = EXTRACT(YEAR FROM NOW()) AND EXTRACT(WEEK FROM dt_create) = EXTRACT (WEEK FROM NOW()) order by dt_create desc "
	rows, err := p.Pool.Query(ctx, sql, userid)
	if err != nil {
		if err.Error() == pgx.ErrNoRows.Error() {
			return nil, models.ErrorNoPosts
		}
		logrus.Errorf("database error with select all posts, %w", err.Error())
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		lastPost := models.LastPost{}
		err = rows.Scan(&lastPost.PostId, &lastPost.Title)
		if err != nil {
			logrus.Errorf("database error with select all posts, %w", err)
			return nil, err
		}
		lastPosts = append(lastPosts, &lastPost)
	}
	return lastPosts, err
}

func (p *DBPostgres) GetAllPosts(ctx context.Context, id string) ([]*models.Post, error) {
	var posts []*models.Post
	sql := "select author.name,author.surname,books.title,posts.content,posts.id from books inner join author on author.id=books.idauthor" +
		" inner join posts on books.id=posts.idbook where posts.userid = $1 or posts.userid in (select unnest(subscriptions) from users where id=$1) order by dt_create desc"
	rows, err := p.Pool.Query(ctx, sql, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, models.ErrorNoPosts
		}
		logrus.Errorf("database error with select all posts, %w", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		po := models.Post{}
		err = rows.Scan(&po.AuthorName, &po.AuthorSurname, &po.Title, &po.Content, &po.PostId)
		if err != nil {
			logrus.Errorf("database error with select all posts, %w", err)
			return nil, err
		}
		posts = append(posts, &po)
	}
	return posts, err
}