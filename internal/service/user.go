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

func (s *Service) GetLastUsers(ctx context.Context) ([]*model.LastUsers, error) {
	return s.rps.GetLastUsersIDs(ctx)
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

func (s *Service) AddSubscriprion(ctx context.Context, subid, id string) error {
	return s.rps.AddSubscriprion(ctx, subid, id)
}

func (s *Service) DeleteSubscription(ctx context.Context, subid, id string) error {
	return s.rps.DeleteSubscription(ctx, subid, id)
}

// DeleteUser delete user from db
func (s *Service) DeleteUser(ctx context.Context, id string) error {
	return s.rps.Delete(ctx, id)
}

func (s *Service) CheckSubs(ctx context.Context, id, userid string) (bool, error) {
	return s.rps.CheckSubs(ctx, id, userid)
}

func (s *Service) GetSubs(ctx context.Context, id string) ([]*model.User, error) {
	return s.rps.GetSubs(ctx, id)
}
