// Package handler ...
package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Egor-Tihonov/Book-network/internal/model"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

var (
	// ErrorStatusUnautharized unutharized
	ErrorStatusUnautharized = errors.New("Unauthorized")
	// tknStr token in string format
	tknStr string
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
		Path:   h.CookiePath,
		Name:   h.CookieName,
		Value:  accessToken,
		MaxAge: h.CookieMaxAge,
	})
	return c.JSON(http.StatusOK, http.NoBody)
}

// Logout ...
func (h *Handler) Logout(c echo.Context) error {
	_, err := h.validation(c)
	c.SetCookie(&http.Cookie{
		Path:   h.CookiePath,
		Name:   h.CookieName,
		Value:  "",
		MaxAge: -1,
	})
	if err != nil {
		if err == ErrorStatusUnautharized {
			return c.JSON(http.StatusUnauthorized, nil)
		}
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, nil)
}

// Validation check token
func (h *Handler) validation(c echo.Context) (*model.JWTClaims, error) {
	cookie, err := c.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			return nil, ErrorStatusUnautharized
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
			return nil, err
		}
		return nil, err
	}
	if !tkn.Valid {
		return nil, err
	}
	return claims, nil
}
