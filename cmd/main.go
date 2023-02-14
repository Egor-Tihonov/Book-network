package main

import (
	"net"

	"github.com/Egor-Tihonov/Book-network/pkg/config"
	"github.com/Egor-Tihonov/Book-network/pkg/pb"
	"github.com/Egor-Tihonov/Book-network/pkg/repository"
	"github.com/Egor-Tihonov/Book-network/pkg/servers"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		logrus.Fatalf("error load configs: %w", err)
	}

	db, err := repository.New(c.DBUrl)
	if err != nil {
		logrus.Fatalf("error connecting to db, %w", err)
	}

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		logrus.Fatalln("Failed to listing:", err)
	}

	logrus.Info("------ START SERVER ON ", c.Port, " ------")

	s := servers.New(db)

	grpcServer := grpc.NewServer()

	pb.RegisterUserServiceServer(grpcServer, s)

	if err := grpcServer.Serve(lis); err != nil {
		logrus.Fatalln("Failed to serve:", err)
	}
}
