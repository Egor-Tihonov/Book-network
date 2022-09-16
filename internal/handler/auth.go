// Package handler ...
package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Egor-Tihonov/Book-network/internal/model"
	"github.com/Egor-Tihonov/Book-network/internal/server"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

// Registration create new user
func (h *Handler) Registration(c echo.Context) error {
	person := model.UserModel{}
	err := json.NewDecoder(c.Request().Body).Decode(&person)
	if err != nil {
		log.Errorf("failed parse json, %e", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	err = h.se.RegistrationUser(c.Request().Context(), &person)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, nil)
}

// Authentication login, create tokens and push it in cookies
func (h *Handler) Authentication(c echo.Context) error {
	authForm := model.AuthenticationForm{}
	err := json.NewDecoder(c.Request().Body).Decode(&authForm)
	if err != nil {
		log.Errorf("failed parse json, %e", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	accessToken, err := h.se.Authentication(c.Request().Context(), &authForm)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	c.SetCookie(&http.Cookie{
		Name:    "token",
		Value:   accessToken,
		Expires: model.ExpirationTime,
	})
	return c.JSON(http.StatusOK, http.NoBody)
}

// Logout ...
func (h *Handler) Logout(c echo.Context) error {
	claims, err := h.se.Validation(c)
	if err != nil {
		if err == server.ErrorStatusUnautharized {
			return c.JSON(http.StatusUnauthorized, nil)
		}
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	claims.ExpiresAt = time.Now().Unix()
	return c.JSON(http.StatusOK, nil)
}
