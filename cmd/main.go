package main

import (
	"fmt"
	"net/http"

	"github.com/Egor-Tihonov/Book-network/internal/config"
	"github.com/Egor-Tihonov/Book-network/internal/handler"
	mid "github.com/Egor-Tihonov/Book-network/internal/middleware"
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

	repo, err := repository.New( /*cfg.PostgresDBURL*/ )
	if err != nil {
		logrus.Fatalf("Connection was failed, %e", err)
	}
	defer repo.Pool.Close()

	srv := service.New(repo, []byte( /*cfg.JWTKey*/ "SUPER-KEY"), model.MyCookie{
		CookieTokenName: "token",
		CookieUserName:  "user", //cfg.CookieName,
		CookiePath:      "/",
	},
	)
	h := handler.New(srv)

	e := echo.New()
	/*g := e.Group("/user")
	g.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:      &model.JWTClaims{},
		SigningKey:  []byte("SUPER-KEY"), //cfg.JWTKey
		TokenLookup: "cookie:token",
		ErrorHandlerWithContext: JWTErrorChecker,
	}))*/

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		AllowCredentials: true,
	}))
	e.POST("/login", h.Authentication)
	e.POST("/sign-up", h.Registration)
	e.GET("/user", h.GetUser, mid.IsLoggedIn)
	e.PUT("/user", h.UpdateUser, mid.IsLoggedIn)
	e.DELETE("/user", h.DeleteUser, mid.IsLoggedIn)
	e.POST("/user/logout", h.Logout)
	e.POST("/user/new-post", h.CreatePost, mid.IsLoggedIn)
	e.GET("/user/posts", h.GetPosts, mid.IsLoggedIn)

	err = e.Start(":8000")
	if err != nil {
		repo.Pool.Close()
		logrus.Fatalf("error started service")
	}
}
