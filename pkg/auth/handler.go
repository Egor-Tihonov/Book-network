package auth

import (
	"github.com/Egor-Tihonov/Book-network/pkg/auth/handlers"
	"github.com/Egor-Tihonov/Book-network/pkg/config"
	"github.com/Egor-Tihonov/Book-network/pkg/models"
	pb "github.com/Egor-Tihonov/Book-network/pkg/pb/auth"
)

func RegisterHandlers(conf config.Config) *ServiceClient {
	svc := ServiceClient{
		Cliet: InitAuthClient(&conf),
	}
	return &svc
}

func (s *ServiceClient) DeleteUser(id string) (*pb.Response, error) {
	return handlers.DeleteUser(id, s.Cliet)
}

func (s *ServiceClient) UpdatePassword(user *models.PasswordUpdate) (*pb.Response, error) {
	return handlers.UpdatePassword(user, s.Cliet)
}
