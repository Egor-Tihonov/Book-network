package services

import (
	"context"

	"github.com/Egor-Tihonov/Book-network/pkg/models"
)

func (s *Service) GetMyFeed(ctx context.Context, id string) ([]*models.Feed, error) {
	return s.rps.GetPosts(ctx, id)
}

// CreateUser: go to repo and add user to db
func (s *Service) CreateUser(ctx context.Context, user *models.UserModel) error {
	return s.rps.Create(ctx, user)
}

// GetUser: get user logic and get user from db by id
func (s *Service) GetUser(ctx context.Context, id string) (*models.User, []*models.Post, []*models.User, error) {
	user, err := s.rps.GetUser(ctx, id)
	if err != nil {
		return nil, nil, nil, err
	}

	posts, err := s.rps.GetMyPosts(ctx, id)
	if err != nil {
		return nil, nil, nil, err
	}

	subs, err := s.rps.GetMySubs(ctx, id)
	if err != nil {
		return nil, nil, nil, err
	}

	return user, posts, subs, nil
}

// DeleteUser: go to repo and add delete user from db
func (s *Service) DeleteUser(ctx context.Context, id string) error {
	return s.rps.Delete(ctx, id)
}

// UpdateUser: go to repo and update user
func (s *Service) UpdateUser(ctx context.Context, id string, user *models.UserUpdate) error {
	if user.Status != "" {
		err := s.rps.Update(ctx, id, user, "status")
		if err != nil {
			return err
		}
	}

	if user.Name != "" {
		err := s.rps.Update(ctx, id, user, "name")
		if err != nil {
			return err
		}
	}

	if user.Username != "" {
		err := s.rps.Update(ctx, id, user, "username")
		if err != nil {
			return err
		}
	}

	return nil
}

// GetNewUsers: go to repo and get ids of new users
func (s *Service) GetNewUsers(ctx context.Context, id string) ([]*models.LastUsers, error) {
	return s.rps.GetLastUsersIDs(ctx, id)
}

func (s *Service) AddNewSubscription(ctx context.Context, subid, id string) error {
	return s.rps.AddSubscriprion(ctx, subid, id)
}

func (s *Service) DeleteOneSubscription(ctx context.Context, subid, id string) error {
	return s.rps.DeleteSubscription(ctx, subid, id)
}

func (s *Service) CheckSubs(ctx context.Context, myid, userid string) (string, error) {
	return s.rps.CheckSubs(ctx, userid, myid)
}

func (s *Service) GetMySubs(ctx context.Context, id string) ([]*models.User, error) {
	return s.rps.GetMySubs(ctx, id)
}

func (s *Service) SearchUser(ctx context.Context, query string) ([]*models.User, error) {
	return s.rps.Search(ctx, query)
}
