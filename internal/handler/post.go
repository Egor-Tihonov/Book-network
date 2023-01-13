package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Egor-Tihonov/Book-network/internal/model"
	"github.com/labstack/echo/v4"
)

func (h *Handler) CreatePost(c echo.Context) error {
	claims, err := h.validation(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}
	post := model.Post{}
	err = json.NewDecoder(c.Request().Body).Decode(&post)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err = h.se.NewPost(c.Request().Context(), claims.ID, &post)
	if err != nil {
		return echo.NewHTTPError(404, err.Error())
	}
	return c.JSON(http.StatusOK, nil)
}
func (h *Handler) GetPosts(c echo.Context) error {
	claims, err := h.validation(c)
	if err != nil {
		if err == echo.ErrUnauthorized {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	posts, err := h.se.GetPosts(c.Request().Context(), claims.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, posts)
}
