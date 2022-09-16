package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Egor-Tihonov/Book-network/internal/model"
	"github.com/Egor-Tihonov/Book-network/internal/server"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

var (
	tknStr string
)

//Handler ...
type Handler struct {
	se *server.Server
}

//Create new handler
func New(srv *server.Server) *Handler {
	return &Handler{se: srv}
}

//GetUser get info about user from db
func (h *Handler) GetUser(c echo.Context) error {
	cookie, err := c.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			return c.JSON(http.StatusUnauthorized, err.Error())
		}
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	tknStr = cookie.Value
	claims, err := validation(tknStr)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err.Error())
	}
	user, err := h.se.GetUser(context.Background(), claims.Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}

func (h *Handler) UpdateUser(c echo.Context) error {
	cookie, err := c.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			return c.JSON(http.StatusUnauthorized, err.Error())
		}
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	tknStr = cookie.Value
	claims, err := validation(tknStr)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err.Error())
	}
	person := model.UserModel{}
	err = json.NewDecoder(c.Request().Body).Decode(&person)
	if err != nil {
		log.Errorf("failed parse json, %e", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	err = h.se.UpdateUser(context.Background(), claims.Id, &person)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, nil)
}
