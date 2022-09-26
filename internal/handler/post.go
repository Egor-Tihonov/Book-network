package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Egor-Tihonov/Book-network/internal/model"
	"github.com/labstack/echo/v4"
)

func (h *Handler) CreatePost(c echo.Context) error {
	claims, err := h.validation(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err.Error())
	}
	post := model.Post{}
	err = json.NewDecoder(c.Request().Body).Decode(&post)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	err = h.se.NewPost(c.Request().Context(), claims.ID, &post)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, nil)
}
func (h *Handler) GetPosts(ctx context.Context, id string) ([]*model.Post, error) {
	return h.se.GetPosts(ctx, id)
}
