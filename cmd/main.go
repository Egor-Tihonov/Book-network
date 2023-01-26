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
		CookieUserName:  "user",
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
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		AllowCredentials: true,
	}))
	e.POST("/login", h.Authentication)
	e.POST("/sign-up", h.Registration)

	e.POST("/user/logout", h.Logout)
	e.PUT("/user/update", h.UpdateUser, mid.IsLoggedIn)
	e.DELETE("/user/delete", h.DeleteUser, mid.IsLoggedIn)
	e.GET("/user/:id", h.GetOtherUser, mid.IsLoggedIn)
	e.GET("/user", h.GetUser, mid.IsLoggedIn)
	e.GET("/user/following", h.MySubscriptions, mid.IsLoggedIn)
	e.GET("/newusers", h.GetLastUsers, mid.IsLoggedIn)
	e.PUT("/user/add-subscription/:id", h.AddSubscription, mid.IsLoggedIn)
	e.PUT("/user/remove-subscription/:id", h.DeleteSubscription, mid.IsLoggedIn)

	e.GET("/:id/posts", h.GetPosts, mid.IsLoggedIn)
	e.GET("/user/last-posts", h.GetLastPosts, mid.IsLoggedIn)
	e.GET("/post/:id", h.GetPost, mid.IsLoggedIn)
	e.POST("/user/new-post", h.CreatePost, mid.IsLoggedIn)
	e.GET("/posts", h.GetAllPosts, mid.IsLoggedIn)

	err = e.Start(":8000")
	if err != nil {
		repo.Pool.Close()
		logrus.Fatalf("error started service")
	}
}
