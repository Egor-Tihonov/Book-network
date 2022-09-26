// Package handler ...
package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Egor-Tihonov/Book-network/internal/model"
	"github.com/Egor-Tihonov/Book-network/internal/server"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

// Handler ...
type Handler struct {
	se *server.Server
	model.MyCookie
}

// New create new handler
func New(srv *server.Server, c model.MyCookie) *Handler {
	return &Handler{se: srv, MyCookie: c}
}

// GetUser get info about user from db
func (h *Handler) GetUser(c echo.Context) error {
	claims, err := h.validation(c)
	if err != nil {
		if err == ErrorStatusUnautharized {
			return c.JSON(http.StatusUnauthorized, err.Error())
		}
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	user, err := h.se.GetUser(c.Request().Context(), claims.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	posts, err := h.GetPosts(c.Request().Context(), claims.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	page := getPage(posts, user)
	return c.JSON(http.StatusOK, page)
}

// UpdateUser update user in db
func (h *Handler) UpdateUser(c echo.Context) error {
	claims, err := h.validation(c)
	if err != nil {
		if err == ErrorStatusUnautharized {
			return c.JSON(http.StatusUnauthorized, err.Error())
		}
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	person := model.UserModel{}
	err = json.NewDecoder(c.Request().Body).Decode(&person)
	if err != nil {
		log.Errorf("failed parse json, %e", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	err = h.se.UpdateUser(c.Request().Context(), claims.ID, &person)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, nil)
}

// DeleteUser delete user from db
func (h *Handler) DeleteUser(c echo.Context) error {
	claims, err := h.validation(c)
	if err != nil {
		if err == ErrorStatusUnautharized {
			return c.JSON(http.StatusUnauthorized, err.Error())
		}
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	err = h.se.DeleteUser(c.Request().Context(), claims.ID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	c.SetCookie(&http.Cookie{
		SameSite: http.SameSiteLaxMode,
		HttpOnly: true,
		Name:     h.CookieName,
		Value:    "",
		MaxAge:   -1,
	})
	if err != nil {
		if err == ErrorStatusUnautharized {
			return c.JSON(http.StatusUnauthorized, nil)
		}
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, nil)
}

func getPage(posts []*model.Post, user *model.UserModel) *model.Page {
	page := &model.Page{
		Username: user.Username,
		Name:     user.Name,
		Posts:    posts,
	}
	return page
}
