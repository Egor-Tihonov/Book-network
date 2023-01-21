// Package handler ...
package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Egor-Tihonov/Book-network/internal/model"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

// GetUser get info about user from db
func (h *Handler) GetUser(c echo.Context) error {
	cookie, err := c.Cookie("user")
	if err != nil {
		return echo.NewHTTPError(404, err)
	}

	user, err := h.se.GetUser(c.Request().Context(), cookie.Value)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

// UpdateUser update user in db
func (h *Handler) UpdateUser(c echo.Context) error {
	cookie, err := c.Cookie("user")
	if err != nil {
		return echo.NewHTTPError(404, err)
	}

	newClaims := model.UserUpdate{}
	err = json.NewDecoder(c.Request().Body).Decode(&newClaims)
	if err != nil {
		log.Errorf("failed parse json, %e", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	err = h.se.UpdateUser(c.Request().Context(), cookie.Value, &newClaims)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, nil)
}

// DeleteUser delete user from db
func (h *Handler) DeleteUser(c echo.Context) error {
	cookie, err := c.Cookie("user")
	if err != nil {
		return echo.NewHTTPError(404, err)
	}
	
	err = h.se.DeleteUser(c.Request().Context(), cookie.Value)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	
	c.SetCookie(&http.Cookie{
		Name:   "user",
		Path:   h.se.Co.CookiePath,
		Value:  "",
		MaxAge: -1,
	})

	c.SetCookie(&http.Cookie{
		Name:   "token",
		Path:   h.se.Co.CookiePath,
		Value:  "",
		MaxAge: -1,
	})

	return c.JSON(http.StatusOK, nil)
}
