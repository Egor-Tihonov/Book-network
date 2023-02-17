package services

import (
	"github.com/Egor-Tihonov/Book-network/pkg/auth"
	"github.com/Egor-Tihonov/Book-network/pkg/db"
)

// Handler ...
type Service struct {
	rps    *db.DBPostgres
	Client *auth.ServiceClient
}

// New create new handler
func New(r *db.DBPostgres, cl *auth.ServiceClient) *Service {
	return &Service{
		rps: r,
		Client: cl,
	}
}
