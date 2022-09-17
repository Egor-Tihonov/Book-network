package main

import (
	"fmt"
	"net/http"

	"github.com/Egor-Tihonov/Book-network/internal/config"
	"github.com/Egor-Tihonov/Book-network/internal/handler"
	"github.com/Egor-Tihonov/Book-network/internal/model"
	"github.com/Egor-Tihonov/Book-network/internal/repository"
	"github.com/Egor-Tihonov/Book-network/internal/server"
	"github.com/caarlos0/env"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg := config.Config{}
	err := env.Parse(&cfg)
	if err != nil {
		logrus.Fatalf("Error parsing env %e", err)
	}
	fmt.Println(cfg)
	repo, err := repository.New(cfg.PostgresDBURL)
	if err != nil {
		logrus.Fatalf("Connection was failed, %e", err)
	}
	defer repo.Pool.Close()
	srv := server.New(repo, []byte(cfg.JWTKey))
	h := handler.New(srv, model.MyCookie{
		CookieName:   cfg.CookieName,
		CookieMaxAge: cfg.CookieMaxAge,
		CookiePath:   "/",
	})

	e := echo.New()
	e.POST("/login", h.Authentication)
	e.POST("/sign-up", h.Registration)
	e.GET("/user", h.GetUser)
	e.PUT("/user/update", h.UpdateUser)
	e.DELETE("/user/delete", h.DeleteUser)
	e.POST("/user/logout", h.Logout)

	err = http.ListenAndServe(":8000", e)
	if err != nil {
		repo.Pool.Close()
		logrus.Fatalf("error started server")
	}
}
