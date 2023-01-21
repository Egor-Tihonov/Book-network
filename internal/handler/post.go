package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Egor-Tihonov/Book-network/internal/model"
	"github.com/labstack/echo/v4"
)

func (h *Handler) CreatePost(c echo.Context) error {
	cookie, err := c.Cookie("user")
	if err != nil {
		return echo.NewHTTPError(401, nil)
	}
	post := model.Post{}
	err = json.NewDecoder(c.Request().Body).Decode(&post)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err = h.se.NewPost(c.Request().Context(), cookie.Value, &post)
	if err != nil {
		return echo.NewHTTPError(404, err.Error())
	}
	return c.JSON(http.StatusOK, nil)
}
func (h *Handler) GetPosts(c echo.Context) error {
	cookie, err := c.Cookie("user")
	if err != nil {
		return echo.NewHTTPError(401, nil)
	}
	posts, err := h.se.GetPosts(c.Request().Context(), cookie.Value)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, posts)
}
