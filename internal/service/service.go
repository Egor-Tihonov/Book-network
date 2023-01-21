// Package service ...
package service

import (
	"github.com/Egor-Tihonov/Book-network/internal/model"
	"github.com/Egor-Tihonov/Book-network/internal/repository"
)

// service ...
type Service struct {
	rps    *repository.PostgresDB
	JWTKey []byte
	Co     *model.MyCookie
}

// New create new connection
func New(repo *repository.PostgresDB, jwtKey []byte, cookie model.MyCookie) *Service {
	return &Service{rps: repo, JWTKey: jwtKey, Co: &cookie}
}
