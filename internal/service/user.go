// Package service ...
package service

import (
	"context"

	"github.com/Egor-Tihonov/Book-network/internal/model"
)

// GetUser get user from db
func (s *Service) GetUser(ctx context.Context, id string) (*model.User, error) {
	return s.rps.Get(ctx, id)
}

// UpdateUser add/replace new information
func (s *Service) UpdateUser(ctx context.Context, id string, mdl *model.UserUpdate) error {
	return s.rps.Update(ctx, id, mdl)
}

// DeleteUser delete user from db
func (s *Service) DeleteUser(ctx context.Context, id string) error {
	return s.rps.Delete(ctx, id)
}
