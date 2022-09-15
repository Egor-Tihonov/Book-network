package handler

import (
	"context"
	"net/http"

	"github.com/Egor-Tihonov/Book-network/internal/server"
	"github.com/labstack/echo"
)

type Response struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Token    string `json:"token"`
}
type Handler struct {
	se *server.Server
}

func New(srv *server.Server) *Handler {
	return &Handler{se: srv}
}

func (h *Handler) GetUser(c echo.Context) error {

	cookie, err := c.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			return c.JSON(http.StatusUnauthorized, err.Error())
		}
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	tknStr := cookie.Value
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
