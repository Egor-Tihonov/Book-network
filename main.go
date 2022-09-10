package main

import (
	"context"
	"html/template"
	"io"

	"github.com/Egor-Tihonov/Book-network/internal/cache"
	"github.com/Egor-Tihonov/Book-network/internal/handlers"
	"github.com/Egor-Tihonov/Book-network/internal/middleware"
	"github.com/Egor-Tihonov/Book-network/internal/model"
	"github.com/Egor-Tihonov/Book-network/internal/repository"
	"github.com/Egor-Tihonov/Book-network/internal/service"

	"github.com/caarlos0/env/v6"
	"github.com/go-redis/redis/v9"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// @title Trainee simple API
// @version 1.0
// @description API server for Trainee

// @host localhost:8000
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @name                       Authorization
// @in                         bearer
var (
	poolP *pgxpool.Pool
	poolM *mongo.Client
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
func main() {
	t := &Template{
		templates: template.Must(template.ParseGlob("public/index.html")),
	}
	cfg := model.Config{}
	err := env.Parse(&cfg)
	if err != nil {
		log.Fatalf("failed to start service, %e", err)
	}
	e := echo.New()
	e.Renderer = t
	rdsClient := redisConnection(&cfg)
	conn := DBConnection(&cfg)
	defer func() {
		err = rdsClient.Close()
		if err != nil {
			log.Errorf("error while closing redis connection - %v", err)
		}
		poolP.Close()
		err = poolM.Disconnect(context.Background())
		if err != nil {
			log.Errorf("error close mongo connection - %e", err)
		}
	}()
	c := cache.New(rdsClient)
	rps := service.New(conn, c)
	h := handlers.New(rps)
	//r:= e.Group("")
	e.POST("/sign-up", h.Registration)
	e.PUT("/usersUpdate/:id", h.UpdateUser, middleware.IsAuthenticated)
	e.DELETE("/usersDelete/:id", h.DeleteUser, middleware.IsAuthenticated)
	e.POST("/login", h.Authentication)
	e.POST("/logout/:id", h.Logout, middleware.IsAuthenticated)
	e.GET("/users/:id", h.GetUserByID, middleware.IsAuthenticated)
	e.GET("/refreshToken", h.RefreshToken, middleware.IsAuthenticated)

	err = e.Start(":8000")

	if err != nil {
		defer log.Fatalf("failed to start service, %e", err)
	}
}

// DBConnection create connection with db
func DBConnection(cfg *model.Config) repository.Repository {
	switch cfg.CurrentDB {
	case "postgres":
		poolP, err := pgxpool.Connect(context.Background() /*cfg.PostgresDbUrl */, "postgresql://postgres:123@localhost:5432/person")
		if err != nil {
			log.Errorf("bad connection with postgresql: %v", err)
			return nil
		}
		return &repository.PRepository{Pool: poolP}

	case "mongo":
		poolM, err := mongo.Connect(context.Background(), options.Client().ApplyURI(cfg.MongoDBURL /*"mongodb://127.0.0.1:27017"*/))
		if err != nil {
			log.Errorf("bad connection with mongoDb: %v", err)
			return nil
		}
		return &repository.MRepository{Pool: poolM}
	}
	return nil
}
func redisConnection(cfg *model.Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", //cfg.RedisURL, /*"localhost:6379"*/
		Password: "",
		DB:       0,
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"status": "connection to redis is failed",
			"err":    err,
		})
		return nil
	}

	logrus.WithFields(logrus.Fields{
		"status": "connection with redis was success",
	})
	return rdb
}
