// Package service ...
package service

import (
	"github.com/Egor-Tihonov/Book-network/internal/repository"
)

// service ...
type Service struct {
	rps    *repository.PostgresDB
	JWTKey []byte
}

// New create new connection
func New(repo *repository.PostgresDB, jwtKey []byte) *Service {
	return &Service{rps: repo, JWTKey: jwtKey}
}
