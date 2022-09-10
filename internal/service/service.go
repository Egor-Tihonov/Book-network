// Package service : file contains server logic
package service

import (
	"context"
	"fmt"

	"github.com/Egor-Tihonov/Book-network/internal/cache"
	"github.com/Egor-Tihonov/Book-network/internal/model"
	"github.com/Egor-Tihonov/Book-network/internal/repository"
)

// JwtKey fo generation and check tokens
var JwtKey = []byte("super-key")

// Service struct
type Service struct {
	rps       repository.Repository
	userCache *cache.UserCache
}

// NewService create new service connection
func New(newRps repository.Repository, userCache *cache.UserCache) *Service { // create
	return &Service{rps: newRps, userCache: userCache}
}

// UpdateUser update user in cache and DB
func (s *Service) UpdateUser(ctx context.Context, id string, person *model.Person) error { // update user
	err := s.rps.Update(ctx, id, person)
	if err != nil {
		return fmt.Errorf("failed to update users, %e", err)
	}
	return s.userCache.AddToCache(ctx, person)
}

// DeleteUser delete user by id from cache
func (s *Service) DeleteUser(ctx context.Context, id string) error { // delete user from DB
	_, found, err := s.userCache.GetUserByIDFromCache(ctx)
	if err != nil {
		return err
	}
	if !found {
		return s.rps.Delete(ctx, id)
	}
	err = s.userCache.DeleteUserFromCache(ctx)
	if err != nil {
		return fmt.Errorf("service: error while deleting user from cache, %e", err)
	}
	return s.rps.Delete(ctx, id)
}

// GetUserByID get user by id from db or cache
func (s *Service) GetUserByID(ctx context.Context, id string) (model.Person, error) { // get one user by id
	user, found, err := s.userCache.GetUserByIDFromCache(ctx)
	if err != nil {
		return model.Person{}, fmt.Errorf("failed to select user from cache, %e", err)
	}
	if !found {
		user, err = s.rps.SelectByID(ctx, id)
		if err != nil {
			return model.Person{}, fmt.Errorf("failed to select user from cache, %e", err)
		}
		err = s.userCache.AddToCache(ctx, &user)
		if err != nil {
			return model.Person{}, fmt.Errorf("failed to select user from cache, %e", err)
		}
		return user, nil
	}
	return user, nil
}

// DeleteFromCache delete user from cache
func (s *Service) DeleteFromCache(ctx context.Context) error {
	_, found, err := s.userCache.GetUserByIDFromCache(ctx)
	if err != nil {
		return err
	}
	if !found {
		return nil
	}
	return s.userCache.DeleteUserFromCache(ctx)
}
