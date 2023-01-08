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
		if err == echo.ErrUnauthorized {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	user, err := h.se.GetUser(c.Request().Context(), claims.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}

// UpdateUser update user in db
func (h *Handler) UpdateUser(c echo.Context) error {
	claims, err := h.validation(c)
	if err != nil {
		if err == echo.ErrUnauthorized {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	newClaims := model.UserUpdate{}
	err = json.NewDecoder(c.Request().Body).Decode(&newClaims)
	if err != nil {
		log.Errorf("failed parse json, %e", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	err = h.se.UpdateUser(c.Request().Context(), claims.ID, &newClaims)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, nil)
}

// DeleteUser delete user from db
func (h *Handler) DeleteUser(c echo.Context) error {
	claims, err := h.validation(c)
	if err != nil {
		if err == echo.ErrUnauthorized {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err = h.se.DeleteUser(c.Request().Context(), claims.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	c.SetCookie(&http.Cookie{
		Name:   h.CookieName,
		Path:   h.CookiePath,
		Value:  "",
		MaxAge: -1,
	})
	if err != nil {
		if err == echo.ErrUnauthorized {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, nil)
}
