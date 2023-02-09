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

// GetUser get info about user from db
func (h *Handler) GetUser(c echo.Context) error {
	userFromJwt := c.Get("user").(*jwt.Token) //why c.Get("user") to get auth header
	claims := userFromJwt.Claims.(jwt.MapClaims)
	idFromParam := claims["id"].(string)

	user, err := h.se.GetUser(c.Request().Context(), idFromParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

func (h *Handler) GetOtherUser(c echo.Context) error {
	id := c.Param("id")

	userFromJwt := c.Get("user").(*jwt.Token) //why c.Get("user") to get auth header
	claims := userFromJwt.Claims.(jwt.MapClaims)
	idFromParam := claims["id"].(string)

	user, err := h.se.GetUser(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	exist, err := h.se.CheckSubs(c.Request().Context(), id, idFromParam)
	if err != nil {
		return echo.NewHTTPError(404, err.Error())
	}

	response := &model.GetOtherUserResponse{
		User:         user,
		Subscription: exist,
	}
	return c.JSON(http.StatusOK, response)
}

func (h *Handler) AddSubscription(c echo.Context) error {
	userFromJwt := c.Get("user").(*jwt.Token) //why c.Get("user") to get auth header
	claims := userFromJwt.Claims.(jwt.MapClaims)
	idFromParam := claims["id"].(string)

	subid := c.Param("id")

	err := h.se.AddSubscriprion(c.Request().Context(), subid, idFromParam)
	if err != nil {
		return echo.NewHTTPError(405, err.Error())
	}
	return c.JSON(200, http.NoBody)

}

func (h *Handler) DeleteSubscription(c echo.Context) error {
	userFromJwt := c.Get("user").(*jwt.Token) //why c.Get("user") to get auth header
	claims := userFromJwt.Claims.(jwt.MapClaims)
	idFromParam := claims["id"].(string)

	subid := c.Param("id")

	err := h.se.DeleteSubscription(c.Request().Context(), subid, idFromParam)
	if err != nil {
		return echo.NewHTTPError(405, err.Error())
	}
	return c.JSON(200, http.NoBody)

}

func (h *Handler) GetLastUsers(c echo.Context) error {
	user, err := h.se.GetLastUsers(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

// UpdateUser update user in db
func (h *Handler) UpdateUser(c echo.Context) error {

	userFromJwt := c.Get("user").(*jwt.Token) //why c.Get("user") to get auth header
	claims := userFromJwt.Claims.(jwt.MapClaims)
	idFromJwt := claims["id"].(string)

	newClaims := model.UserUpdate{}
	err := json.NewDecoder(c.Request().Body).Decode(&newClaims)
	if err != nil {
		log.Errorf("failed parse json, %w", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	err = h.se.UpdateUser(c.Request().Context(), idFromJwt, &newClaims)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, nil)
}

func (h *Handler) GetReviewFeed(c echo.Context) error {
	return nil
}

func (h *Handler) MySubscriptions(c echo.Context) error {
	userFromJwt := c.Get("user").(*jwt.Token) //why c.Get("user") to get auth header
	claims := userFromJwt.Claims.(jwt.MapClaims)
	idFromJwt := claims["id"].(string)

	users, err := h.se.GetSubs(c.Request().Context(), idFromJwt)
	if err != nil {
		return echo.NewHTTPError(404, err.Error())
	}

	return c.JSON(200, users)
}

// DeleteUser delete user from db
func (h *Handler) DeleteUser(c echo.Context) error {
	userFromJwt := c.Get("user").(*jwt.Token) //why c.Get("user") to get auth header
	claims := userFromJwt.Claims.(jwt.MapClaims)
	idFromJwt := claims["id"].(string)

	err := h.se.DeleteUser(c.Request().Context(), idFromJwt)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	c.SetCookie(&http.Cookie{
		Name:   "token",
		Path:   h.se.Co.CookiePath,
		Value:  "",
		MaxAge: -1,
	})

	return c.JSON(http.StatusOK, nil)
}
