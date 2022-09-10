// Package handlers : file contains operation with requests
package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Egor-Tihonov/Book-network/internal/model"
	"github.com/Egor-Tihonov/Book-network/internal/service"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

var validate = validator.New()

// Handler struct
type Handler struct {
	s *service.Service
}

// NewHandler :define new handlers
func New(newS *service.Service) *Handler {
	return &Handler{s: newS}
}

// UpdateUser godoc
// @Summary     UpdateUser
// @Description UpdateUser is echo handler which delete user from cache and db
// @Param       id  path string true "Account ID"
// @Produce     string
// @Tags        User
// @Router      /users/{id} [delete]
// @Failure     500 string
// @Success     200 string
func (h *Handler) UpdateUser(c echo.Context) error {
	person := model.Person{}
	id := c.Param("id")
	err := ValidateValueID(id)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	err = json.NewDecoder(c.Request().Body).Decode(&person)
	if err != nil {
		log.Errorf("failed parse json, %e", err)
		return err
	}
	err = h.s.UpdateUser(c.Request().Context(), id, &person)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusOK, "Ok")
}

// DeleteUser godoc
// @Summary     DeleteUser
// @Description DeleteUser is echo handler which delete user from cache and db
// @Param       id path string true "Account ID"
// @Produce     string
// @Tags        User
// @Router      /users/{id} [delete]
// @Failure     500 json
// @Success     200 string
func (h *Handler) DeleteUser(c echo.Context) error {
	id := c.Param("id")
	err := ValidateValueID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	err = h.s.DeleteUser(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.String(http.StatusOK, "delete")
}

// GetUserByID godoc
// @Summary     GetUserByID
// @Description GetUserByID is echo handler which returns json structure of User object
// @Produce     json
// @Tags        User
// @Param       id path string true "Account ID"
// @Success     200 json
// @Failure     500 json
// @Router      /users/{id} [get]
// @Security    ApiKeyAuth
func (h *Handler) GetUserByID(c echo.Context) error {
	id := c.Param("id")
	err := ValidateValueID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	person, err := h.s.GetUserByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, person)
}

// ValidateValueID validate id
func ValidateValueID(id string) error {
	err := validate.Var(id, "required")
	if err != nil {
		return fmt.Errorf("id length couldnt be less then 36,~%v", err)
	}
	return nil
}
