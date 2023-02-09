package main

import (
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
		logrus.Fatalf("Error parsing env %w", err)
	}
	logrus.Info(cfg)

	repo, err := repository.New( /*cfg.PostgresDBURL*/ )
	if err != nil {
		logrus.Fatalf("Connection was failed, %w", err)
	}
	defer repo.Pool.Close()

	srv := service.New(repo, []byte( /*cfg.JWTKey*/ "SUPER-KEY"), model.MyCookie{
		CookieTokenName: "token",
		CookiePath:      "/",
	},
	)
	h := handler.New(srv)

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		AllowCredentials: true,
	}))
	e.Use(mid.IsLoggedIn)

	//auth handlers
	e.POST("/login", h.Authentication)
	e.POST("/sign-up", h.Registration)
	e.POST("/user/logout", h.Logout)

	//user handlers
	e.PUT("/user/update", h.UpdateUser)
	e.DELETE("/user/delete", h.DeleteUser)
	e.GET("/user/:id", h.GetOtherUser)
	e.GET("/user", h.GetUser)
	e.GET("/user/following", h.MySubscriptions)
	e.GET("/newusers", h.GetLastUsers)
	e.PUT("/user/add-subscription/:id", h.AddSubscription)
	e.PUT("/user/remove-subscription/:id", h.DeleteSubscription)

	//posts handlers
	e.GET("/:id/posts", h.GetPosts)
	e.GET("/user/posts", h.GetMyPosts)
	e.GET("/user/last-posts", h.GetLastPosts)
	e.GET("/post/:id", h.GetPost)
	e.POST("/user/new-post", h.CreatePost)
	e.GET("/posts", h.GetAllPosts)

	err = e.Start(":8000")
	if err != nil {
		repo.Pool.Close()
		logrus.Fatalf("error started service")
	}
}
