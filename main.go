package main

import (
	_ "fmt"

	_ "github.com/Egor-Tihonov/Book-network/internal/config"
	"github.com/Egor-Tihonov/Book-network/internal/handler"
	"github.com/Egor-Tihonov/Book-network/internal/model"
	"github.com/Egor-Tihonov/Book-network/internal/repository"
	"github.com/Egor-Tihonov/Book-network/internal/server"
	_ "github.com/caarlos0/env"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func main() {
	/*cfg := config.Config{}
	err := env.Parse(&cfg)
	if err != nil {
		logrus.Fatalf("Error parsing env %e", err)
	}
	fmt.Println(cfg)*/

	repo, err := repository.New( /*cfg.PostgresDBURL*/ )
	if err != nil {
		logrus.Fatalf("Connection was failed, %e", err)
	}
	defer repo.Pool.Close()

	srv := server.New(repo, []byte("SUPER-KEY" /*cfg.JWTKey*/))
	h := handler.New(srv, model.MyCookie{
		CookieName:   "token", //cfg.CookieName,
		CookieMaxAge: 3600,    //cfg.CookieMaxAge,
		CookiePath:   "/",
	})

	e := echo.New()
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
		logrus.Fatalf("error started server")
	}
}
