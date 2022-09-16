package server

import (
	"context"

	"github.com/Egor-Tihonov/Book-network/internal/model"
	"github.com/Egor-Tihonov/Book-network/internal/repository"
)

type Server struct {
	rps *repository.PostgresDB
}

func New(repo *repository.PostgresDB) *Server {
	return &Server{rps: repo}
}

func (s *Server) GetUser(ctx context.Context, id string) (*model.UserModel, error) {
	return s.rps.SelectUser(ctx, id)
}

func (s *Server) UpdateUser(ctx context.Context, id string, mdl *model.UserModel) error {
	return s.rps.Update(ctx, id, mdl)
}

func (s *Server) Delete(ctx context.Context, id string) error {
	return s.rps.Delete(ctx, id)
}
