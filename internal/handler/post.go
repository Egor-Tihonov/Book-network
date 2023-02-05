package handler

import (
	"encoding/json"
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
		logrus.Errorf("service: error parse json: %e", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err = h.se.NewPost(c.Request().Context(), idFromJwt, &post)
	if err != nil {
		logrus.Errorf("service: create post: %e", err)
		return echo.NewHTTPError(404, err.Error())
	}
	return c.JSON(http.StatusOK, nil)
}

func (h *Handler) GetPosts(c echo.Context) error {
	id := c.Param("id")
	posts, err := h.se.GetPosts(c.Request().Context(), id)
	if err != nil {
		logrus.Errorf("service: get posts: %e", err)
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
		logrus.Errorf("service: get posts: %e", err)
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
		if err == model.ErrorNoPosts {
			return echo.NewHTTPError(404, err.Error())
		}
		logrus.Errorf("error get post: %e", err)
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
		if err == model.ErrorNoPosts {
			return echo.NewHTTPError(404, err.Error())
		}
		logrus.Errorf("error get last post: %e", err)
		return echo.NewHTTPError(405, err.Error())
	}

	return c.JSON(200, lastPosts)
}

func (h *Handler) GetAllPosts(c echo.Context) error {
	userFromJwt := c.Get("user").(*jwt.Token) //why c.Get("user") to get auth header
	claims := userFromJwt.Claims.(jwt.MapClaims)
	idFromJwt := claims["id"].(string)

	posts, err := h.se.GetAllPosts(c.Request().Context(),idFromJwt)
	if err != nil {
		return echo.NewHTTPError(404, err.Error())
	}
	return c.JSON(200, posts)
}
