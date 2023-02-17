package auth

import (
	"github.com/Egor-Tihonov/Book-network/pkg/config"
	pb "github.com/Egor-Tihonov/Book-network/pkg/pb/auth"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Cliet pb.AuthServiceClient
}

func InitAuthClient(conf *config.Config) pb.AuthServiceClient {
	cc, err := grpc.Dial(conf.AuthService, grpc.WithInsecure())
	if err != nil {
		logrus.Errorf("Could not connect: %w", err)
	}

	return pb.NewAuthServiceClient(cc)
}
