package main

import (
	"fmt"
	"net"
	"runtime"
	"strings"

	authSe "github.com/Egor-Tihonov/Book-network/pkg/auth"
	bookSe "github.com/Egor-Tihonov/Book-network/pkg/book-service"
	"github.com/Egor-Tihonov/Book-network/pkg/config"
	"github.com/Egor-Tihonov/Book-network/pkg/db"
	"github.com/Egor-Tihonov/Book-network/pkg/handlers"
	pb "github.com/Egor-Tihonov/Book-network/pkg/pb/user"
	"github.com/Egor-Tihonov/Book-network/pkg/services"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	InitLog()
	c, err := config.LoadConfig()

	svc_a := authSe.RegisterHandlers(c)
	svc_b := bookSe.RegisterHandlers(c)

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

	s := services.New(db, svc_a, svc_b)
	h := handlers.New(s)

	grpcServer := grpc.NewServer()

	pb.RegisterUserServiceServer(grpcServer, h)

	if err := grpcServer.Serve(lis); err != nil {
		logrus.Fatalln("Failed to serve:", err)
	}
}

func InitLog() {
	logrus.SetReportCaller(true)

	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
		ForceColors:     true,
		DisableColors:   false,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			return "", fmt.Sprintf(" %s:%d", formatFilePath(f.File), f.Line)
		},
	})

}

func formatFilePath(path string) string {
	arr := strings.Split(path, "/")
	return arr[len(arr)-1]
}
