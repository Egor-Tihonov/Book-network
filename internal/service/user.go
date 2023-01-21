// Package service ...
package service

import (
	"context"

	"github.com/Egor-Tihonov/Book-network/internal/model"
)

// GetUser get user from db
func (s *Service) GetUser(ctx context.Context, email string) (*model.User, error) {
	return s.rps.Get(ctx, email)
}

// UpdateUser add/replace new information
func (s *Service) UpdateUser(ctx context.Context, id string, user *model.UserUpdate) error {
	oldUser, err := s.rps.GetUserForUpdate(ctx, id)

	if err != nil {
		return err
	}

	if user.Name == "" {
		user.Name = oldUser.Name
	}
	
	if user.Password == "" {
		user.Password = oldUser.Password
	} else {
		newHashPassword, err := s.hashPassword(user.Password)

		if err != nil {
			return err
		}

		user.Password = newHashPassword
	}
	if user.Username == "" {
		user.Username = oldUser.Username
	}

	if user.Status == "" {
		user.Status = oldUser.Status
	}

	return s.rps.Update(ctx, id, user)
}

// DeleteUser delete user from db
func (s *Service) DeleteUser(ctx context.Context, id string) error {
	return s.rps.Delete(ctx, id)
}
