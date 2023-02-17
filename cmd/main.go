package main

import (
	"net"

	authSe "github.com/Egor-Tihonov/Book-network/pkg/auth"
	"github.com/Egor-Tihonov/Book-network/pkg/config"
	"github.com/Egor-Tihonov/Book-network/pkg/db"
	"github.com/Egor-Tihonov/Book-network/pkg/handlers"
	pb "github.com/Egor-Tihonov/Book-network/pkg/pb/user"
	"github.com/Egor-Tihonov/Book-network/pkg/services"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()

	svc := authSe.RegisterHandlers(c)

	if err != nil {
		logrus.Fatalf("error load configs: %w", err)
	}

	db, err := db.New(c.DBUrl)
	if err != nil {
		logrus.Fatalf("error connecting to db, %w", err)
	}

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		logrus.Fatalln("Failed to listing:", err)
	}

	logrus.Info("------ START SERVER ON ", c.Port, " ------")

	s := services.New(db, svc)
	h := handlers.New(s)

	grpcServer := grpc.NewServer()

	pb.RegisterUserServiceServer(grpcServer, h)

	if err := grpcServer.Serve(lis); err != nil {
		logrus.Fatalln("Failed to serve:", err)
	}
}
