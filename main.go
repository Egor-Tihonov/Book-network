package main

import (
	"github.com/Egor-Tihonov/Book-network/internal/handler"
	"github.com/Egor-Tihonov/Book-network/internal/repository"
	"github.com/Egor-Tihonov/Book-network/internal/server"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

func main() {
	repo, err := repository.New("postgresql://postgres:123@localhost:5432/person")
	if err != nil {
		logrus.Fatalf("Connection was failed, %e", err)
	}
	srv := server.New(repo)
	h := handler.New(srv)
	e := echo.New()
	e.POST("/login", h.Authentication)
	e.POST("/sign-up", h.Registration)
	e.GET("/user", h.GetUser)
	e.PUT("/user/update", h.UpdateUser)
	e.DELETE("/user/delete", h.DeleteUser)
	e.POST("/user/logout", h.Logout)
	err = e.Start(":8000")
	if err != nil {
		logrus.Fatalf("error started server")
	}
}
