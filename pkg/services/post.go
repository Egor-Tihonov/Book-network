package services

import (
	"context"
	"errors"
	"time"

	"github.com/Egor-Tihonov/Book-network/pkg/models"
	"github.com/sirupsen/logrus"
)

func (s *Service) CreatePost(ctx context.Context, post *models.Post, id string) error {
	bookid, err := s.Check(ctx, post)
	if err != nil {
		logrus.Errorf("user service: error create post, %w", err.Error())
	}

	bookids, err := s.rps.GetForCheckPosts(ctx, id)
	if err != nil {
		return err
	}

	for _, id := range bookids {
		if id == bookid {
			ErrPostAlreadyExist := errors.New("отзыв на эту книгу уже сотавлен")
			return ErrPostAlreadyExist
		}
	}

	createDate := time.Now()
	return s.rps.CreatePost(ctx, id, bookid, post, createDate)
}

func (s *Service) GetPost(ctx context.Context, postid string) (*models.Post, error) {
	return s.rps.GetPost(ctx, postid)
}

func (s *Service) Check(ctx context.Context, post *models.Post) (string, error) {
	res, err := s.BookClient.GetBook(&models.Book{
		Author: models.Author{
			Name:    post.AuthorName,
			Surname: post.AuthorSurname,
		},
		Title: post.Title,
	})
	logrus.Errorf("%s", err)

	if err != nil {
		return "", err
	}

	return res.Id, nil
}

func (s *Service) GetPosts(ctx context.Context, id string) ([]*models.Feed, error) {
	return s.rps.GetPostsForBook(ctx, id)
}

func (s *Service) DeletePost(ctx context.Context, postid, userid string) error {
	return s.rps.DeletePost(ctx, postid, userid)
}

func (s *Service) GetReviewsByAllUsers(ctx context.Context) ([]*models.Feed, error) {
	return s.rps.GetAllReviewsFromDB(ctx)
}
