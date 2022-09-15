package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Egor-Tihonov/Book-network/internal/model"
	"github.com/Egor-Tihonov/Book-network/internal/server"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func (h *Handler) Registration(c echo.Context) error {
	person := model.UserModel{}
	err := json.NewDecoder(c.Request().Body).Decode(&person)
	if err != nil {
		log.Errorf("failed parse json, %e", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	err = h.se.Registration(context.Background(), &person)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	return c.JSON(http.StatusOK, nil)
}

func (h *Handler) Authentcation(c echo.Context) error {
	authForm := model.AuthentcationForm{}
	err := json.NewDecoder(c.Request().Body).Decode(&authForm)
	if err != nil {
		log.Errorf("failed parse json, %e", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	accessToken, err := h.se.Authentcation(context.Background(), &authForm)
	// resp := &Response{
	// 	Id:       user.Id,
	// 	Username: user.Username,
	// 	Token:    accessToken,
	// }
	if err != nil {
		return c.JSON(400, err.Error())
	}
	c.SetCookie(&http.Cookie{
		Name:    "token",
		Value:   accessToken,
		Expires: model.ExpirationTime,
	})
	return c.JSON(http.StatusOK, http.NoBody)
}

func validation(tknStr string) (model.JWTClaims, error) {
	claims := &model.JWTClaims{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return server.JwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return model.JWTClaims{}, err
		}
		return model.JWTClaims{}, err
	}
	if !tkn.Valid {
		return model.JWTClaims{}, err
	}
	return *claims, nil
}
