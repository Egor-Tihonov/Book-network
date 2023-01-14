package main

import (
	"fmt"
	"net/http"

	"github.com/Egor-Tihonov/Book-network/internal/config"
	"github.com/Egor-Tihonov/Book-network/internal/handler"
	"github.com/Egor-Tihonov/Book-network/internal/model"
	"github.com/Egor-Tihonov/Book-network/internal/repository"
	"github.com/Egor-Tihonov/Book-network/internal/service"
	"github.com/caarlos0/env"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg := config.Config{}
	err := env.Parse(&cfg)
	if err != nil {
		logrus.Fatalf("Error parsing env %e", err)
	}
	fmt.Println(cfg)

	repo, err := repository.New(/*cfg.PostgresDBURL*/)
	if err != nil {
		logrus.Fatalf("Connection was failed, %e", err)
	}
	defer repo.Pool.Close()

	srv := service.New(repo, []byte(/*cfg.JWTKey*/"SUPER-KEY"))
	h := handler.New(srv, model.MyCookie{
		CookieName:   "token", //cfg.CookieName,
		CookieMaxAge: 3600,    //cfg.CookieMaxAge,
		CookiePath:   "/",
	})

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", "https://labstack.net"},
		AllowMethods: []string{http.MethodGet, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))
	e.POST("/login", h.Authentication)
	e.POST("/sign-up", h.Registration)
	e.GET("/user", h.GetUser)
	e.PUT("/user", h.UpdateUser)
	e.DELETE("/user", h.DeleteUser)
	e.POST("/user/logout", h.Logout)
	e.POST("/user/new-post", h.CreatePost)
	e.GET("/user/posts", h.GetPosts)

	err = e.Start(":8000")
	if err != nil {
		repo.Pool.Close()
		logrus.Fatalf("error started service")
	}
}
