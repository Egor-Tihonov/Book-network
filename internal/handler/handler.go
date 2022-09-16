// Package handler ...
package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Egor-Tihonov/Book-network/internal/model"
	"github.com/Egor-Tihonov/Book-network/internal/server"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

// Handler ...
type Handler struct {
	se *server.Server
}

// New create new handler
func New(srv *server.Server) *Handler {
	return &Handler{se: srv}
}

// GetUser get info about user from db
func (h *Handler) GetUser(c echo.Context) error {
	claims, err := h.se.Validation(c)
	if err != nil {
		if err == server.ErrorStatusUnautharized {
			return c.JSON(http.StatusUnauthorized, nil)
		}
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	user, err := h.se.GetUser(c.Request().Context(), claims.Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}

// UpdateUser update user in db
func (h *Handler) UpdateUser(c echo.Context) error {
	claims, err := h.se.Validation(c)
	if err != nil {
		if err == server.ErrorStatusUnautharized {
			return c.JSON(http.StatusUnauthorized, nil)
		}
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	person := model.UserModel{}
	err = json.NewDecoder(c.Request().Body).Decode(&person)
	if err != nil {
		log.Errorf("failed parse json, %e", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	err = h.se.UpdateUser(c.Request().Context(), claims.Id, &person)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, nil)
}

// DeleteUser delete user from db
func (h *Handler) DeleteUser(c echo.Context) error {
	claims, err := h.se.Validation(c)
	if err != nil {
		if err == server.ErrorStatusUnautharized {
			return c.JSON(http.StatusUnauthorized, nil)
		}
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	err = h.se.DeleteUser(c.Request().Context(), claims.Id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, nil)
}
