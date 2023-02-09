package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Egor-Tihonov/Book-network/internal/model"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func (h *Handler) CreatePost(c echo.Context) error {
	userFromJwt := c.Get("user").(*jwt.Token) //why c.Get("user") to get auth header
	claims := userFromJwt.Claims.(jwt.MapClaims)
	idFromJwt := claims["id"].(string)

	post := model.Post{}
	err := json.NewDecoder(c.Request().Body).Decode(&post)
	if err != nil {
		logrus.Errorf("error parse json: %w", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err = h.se.NewPost(c.Request().Context(), idFromJwt, &post)
	if err != nil {
		logrus.Errorf("create post error: %w", err)
		return echo.NewHTTPError(404, err.Error())
	}
	return c.JSON(http.StatusOK, nil)
}

func (h *Handler) GetPosts(c echo.Context) error {
	id := c.Param("id")
	posts, err := h.se.GetPosts(c.Request().Context(), id)
	if err != nil {
		logrus.Errorf("get all user posts error: %w", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, posts)
}

func (h *Handler) GetMyPosts(c echo.Context) error {
	userFromJwt := c.Get("user").(*jwt.Token) //why c.Get("user") to get auth header
	claims := userFromJwt.Claims.(jwt.MapClaims)
	idFromJwt := claims["id"].(string)

	posts, err := h.se.GetPosts(c.Request().Context(), idFromJwt)
	if err != nil {
		logrus.Errorf("get my posts error: %w", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, posts)
}

func (h *Handler) GetPost(c echo.Context) error {
	userFromJwt := c.Get("user").(*jwt.Token) //why c.Get("user") to get auth header
	claims := userFromJwt.Claims.(jwt.MapClaims)
	idFromJwt := claims["id"].(string)

	postID := c.Param("id")

	post, err := h.se.GetPost(c.Request().Context(), idFromJwt, postID)
	if err != nil {
		if errors.Is(err, model.ErrorNoPosts) {
			return echo.NewHTTPError(404, err.Error())
		}
		logrus.Errorf("get one post error: %w", err)
		return echo.NewHTTPError(405, err.Error())
	}

	return c.JSON(http.StatusOK, post)
}

func (h *Handler) GetLastPosts(c echo.Context) error {
	userFromJwt := c.Get("user").(*jwt.Token) //why c.Get("user") to get auth header
	claims := userFromJwt.Claims.(jwt.MapClaims)
	idFromJwt := claims["id"].(string)

	lastPosts, err := h.se.GetLast(c.Request().Context(), idFromJwt)
	if err != nil {
		if errors.Is(err, model.ErrorNoPosts) {
			return echo.NewHTTPError(404, err.Error())
		}
		logrus.Errorf("get last post error: %w", err)
		return echo.NewHTTPError(405, err.Error())
	}

	return c.JSON(200, lastPosts)
}

func (h *Handler) GetAllPosts(c echo.Context) error {
	userFromJwt := c.Get("user").(*jwt.Token) //why c.Get("user") to get auth header
	claims := userFromJwt.Claims.(jwt.MapClaims)
	idFromJwt := claims["id"].(string)

	posts, err := h.se.GetAllPosts(c.Request().Context(), idFromJwt)
	if err != nil {
		logrus.Errorf("get my feed posts error: %w", err)
		return echo.NewHTTPError(404, err.Error())
	}
	return c.JSON(200, posts)
}
