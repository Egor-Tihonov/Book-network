// Package handler ...
package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Egor-Tihonov/Book-network/internal/model"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

// Registration create new user
func (h *Handler) Registration(c echo.Context) error {
	user := model.UserModel{}
	err := json.NewDecoder(c.Request().Body).Decode(&user)
	if err != nil {
		log.Errorf("handler: failed parse json, %e", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	err = h.se.RegistrationUser(c.Request().Context(), &user)
	if err != nil {
		return echo.NewHTTPError(404, err.Error())
	}
	return c.JSON(http.StatusOK, nil)
}

// Authentication login, create tokens and push it in cookies
func (h *Handler) Authentication(c echo.Context) error {
	authForm := model.AuthenticationForm{}
	err := json.NewDecoder(c.Request().Body).Decode(&authForm)
	if err != nil {
		log.Errorf("handler: failed parse json, %e", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	user, err := h.se.Authentication(c.Request().Context(), &authForm)
	if err != nil {
		log.Errorf("handler: failed with auth, %e", err)
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	cookieToken, cookieUser, err := h.se.GenerateTokensAndSetCookies(user, c.Request().Context())

	if err != nil {
		log.Errorf("handler: failed with auth, %e", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	c.SetCookie(cookieToken)
	c.SetCookie(cookieUser)

	return c.JSON(http.StatusOK, http.NoBody)
}

// Logout ...
func (h *Handler) Logout(c echo.Context) error {
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
