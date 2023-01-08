package server

import (
	"context"
	"errors"

	"github.com/Egor-Tihonov/Book-network/internal/model"
)

func (s *Server) NewPost(ctx context.Context, userId string, post *model.Post) error {
	if(post.Content=="" || post.Title == "" || post.AuthorName == "" || post.AuthorSurname == ""){
		return nil
	}
	ids, err := s.rps.GetForCheckPosts(ctx, userId)
	if err != nil {
		return err
	}
	bookId, authorId, err := s.rps.GetForCheckAuthor(ctx, post.AuthorName, post.AuthorSurname)
	for _, i := range ids {
		if i == bookId {
			err = errors.New("you already create post for this author and book")
			return err
		}
	}
	if err != nil {
		return err
	}
	err = s.rps.CreatePost(ctx, userId, authorId, bookId, post)
	return err
}
func (s *Server) GetPosts(ctx context.Context, id string) ([]*model.Post, error) {
	return s.rps.GetAll(ctx, id)
}