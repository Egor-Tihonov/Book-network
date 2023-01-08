// Package server ...
package server

import (
	"context"

	"github.com/Egor-Tihonov/Book-network/internal/model"
	"github.com/Egor-Tihonov/Book-network/internal/repository"
)

// Server ...
type Server struct {
	rps    *repository.PostgresDB
	JWTKey []byte
}

// New create new connection
func New(repo *repository.PostgresDB, jwtKey []byte) *Server {
	return &Server{rps: repo}
}

// GetUser get user from db
func (s *Server) GetUser(ctx context.Context, id string) (*model.User, error) {
	return s.rps.Get(ctx, id)
}

// UpdateUser add/replace new information
func (s *Server) UpdateUser(ctx context.Context, id string, mdl *model.UserUpdate) error {
	return s.rps.Update(ctx, id, mdl)
}

// DeleteUser delete user from db
func (s *Server) DeleteUser(ctx context.Context, id string) error {
	return s.rps.Delete(ctx, id)
}
