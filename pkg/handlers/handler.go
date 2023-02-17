package handlers

import (
	pb "github.com/Egor-Tihonov/Book-network/pkg/pb/user"
	"github.com/Egor-Tihonov/Book-network/pkg/services"
)

type Handler struct {
	se *services.Service
	pb.UnimplementedUserServiceServer
}

func New(s *services.Service) *Handler {
	return &Handler{
		se: s,
	}
}
