package bookservice

import (
	"github.com/Egor-Tihonov/Book-network/pkg/config"
	pb "github.com/Egor-Tihonov/Book-network/pkg/pb/book"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.BookServiceClient
}

func InitBookClient(conf *config.Config) pb.BookServiceClient {
	cc, err := grpc.Dial(conf.BookService, grpc.WithInsecure())
	if err != nil {
		logrus.Errorf("Could not connect: %w", err)
	}

	return pb.NewBookServiceClient(cc)
}
