package service

import (
	"context"
	"errors"
	"time"

	"github.com/Egor-Tihonov/Book-network/internal/model"
)

//NewPost: create new post and check for error(people cant create the same review 2 times)
func (s *Service) NewPost(ctx context.Context, userId string, post *model.Post) error {
	if post.Content == "" || post.Title == "" || post.AuthorName == "" || post.AuthorSurname == "" {
		return nil
	}

	//get id`s of books with user review
	ids, err := s.rps.GetForCheckPosts(ctx, userId)
	if err != nil {
		return err
	}

	//get bookId and auhro id with the same names
	bookId, authorId, err := s.rps.GetForCheckAuthor(ctx, post.AuthorName, post.AuthorSurname)
	if err != nil {
		return err
	}

	//check id`s
	for _, i := range ids {
		if i == bookId {
			err = errors.New("you already create post for this author and book")
			return err
		}
	}

	err = s.rps.CreatePost(ctx, userId, authorId, bookId, post, time.Now())
	return err
}
func (s *Service) GetPosts(ctx context.Context, id string) ([]*model.Post, error) {
	return s.rps.GetAll(ctx, id)
}

func (s *Service) GetPost(ctx context.Context, userid, postid string) (*model.Post, error) {
	return s.rps.GetPost(ctx, userid, postid)
}

func (s *Service) GetLast(ctx context.Context, userid string) ([]*model.LastPost, error) {
	return s.rps.GetLast(ctx, userid)
}

func (s *Service) GetAllPosts(ctx context.Context, id string) ([]*model.Post, error) {
	return s.rps.GetAllPosts(ctx,id)
}

