package services

import (
	"github.com/Egor-Tihonov/Book-network/pkg/auth"
	bookservice "github.com/Egor-Tihonov/Book-network/pkg/book-service"
	"github.com/Egor-Tihonov/Book-network/pkg/db"
)

// Handler ...
type Service struct {
	rps        *db.DBPostgres
	AuthClient *auth.ServiceClient
	BookClient *bookservice.ServiceClient
}

// New create new handler
func New(r *db.DBPostgres, cl_a *auth.ServiceClient, cl_b *bookservice.ServiceClient) *Service {
	return &Service{
		rps:        r,
		AuthClient: cl_a,
		BookClient: cl_b,
	}
}
