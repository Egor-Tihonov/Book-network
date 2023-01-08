// Package handler ...
package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Egor-Tihonov/Book-network/internal/model"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

var (

	// tknStr token in string format
	tknStr string
)

// Registration create new user
func (h *Handler) Registration(c echo.Context) error {
	person := model.UserModel{}
	err := json.NewDecoder(c.Request().Body).Decode(&person)
	if err != nil {
		log.Errorf("failed parse json, %e", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	err = h.se.RegistrationUser(c.Request().Context(), &person)
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
		log.Errorf("failed parse json, %e", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	accessToken, err := h.se.Authentication(c.Request().Context(), &authForm)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	c.SetCookie(&http.Cookie{
		Name:  h.CookieName,
		Path:  h.CookiePath,
		Value: accessToken,
	})
	return c.JSON(http.StatusOK, http.NoBody)
}

// Logout ...
func (h *Handler) Logout(c echo.Context) error {
	c.SetCookie(&http.Cookie{
		Name:   h.CookieName,
		Path:   h.CookiePath,
		Value:  "",
		MaxAge: -1,
	})
	return c.JSON(http.StatusOK, nil)
}

// Validation check token
func (h *Handler) validation(c echo.Context) (*model.JWTClaims, error) {
	cookie, err := c.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			return nil, echo.ErrUnauthorized
		}
		return nil, err
	}
	tknStr = cookie.Value
	claims := &model.JWTClaims{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return h.se.JWTKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, echo.ErrUnauthorized
		}
		return nil, err
	}
	if !tkn.Valid {
		return nil, echo.ErrUnauthorized
	}
	return claims, nil
}
