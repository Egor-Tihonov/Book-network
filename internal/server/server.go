package server

import (
	"context"

	"github.com/Egor-Tihonov/Book-network/internal/model"
	"github.com/Egor-Tihonov/Book-network/internal/repository"
)

//Server ...
type Server struct {
	rps *repository.PostgresDB
}

//New create new connection
func New(repo *repository.PostgresDB) *Server {
	return &Server{rps: repo}
}

//GetUser get user from db
func (s *Server) GetUser(ctx context.Context, id string) (*model.UserModel, error) {
	return s.rps.Get(ctx, id)
}

//UpdateUser add/replace new information
func (s *Server) UpdateUser(ctx context.Context, id string, mdl *model.UserModel) error {
	return s.rps.Update(ctx, id, mdl)
}

//Delete delete user from db
func (s *Server) DeleteUser(ctx context.Context, id string) error {
	return s.rps.Delete(ctx, id)
}
