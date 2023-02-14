package servers

import (
	"github.com/Egor-Tihonov/Book-network/pkg/pb"
	"github.com/Egor-Tihonov/Book-network/pkg/repository"
)

// Handler ...
type Server struct {
	pb.UnimplementedUserServiceServer
	Rps *repository.DBPostgres
}

// New create new handler
func New(r *repository.DBPostgres) *Server {
	return &Server{Rps: r}
}
