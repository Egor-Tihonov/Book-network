package services

import (
	"context"

	"github.com/Egor-Tihonov/Book-network/pkg/models"
)

// GetUser get info about user from db
func (s *Service) CreateUser(ctx context.Context, user *models.UserModel) error {
	return s.rps.Create(ctx, user)
}

func (s *Service) GetUser(ctx context.Context, id string) (*models.User, []*models.Post, []*models.User, error) {
	user, err := s.rps.GetUser(ctx, id)
	if err != nil {
		return nil, nil, nil, err
	}

	posts, err := s.rps.GetAllPosts(ctx, id)
	if err != nil {
		return nil, nil, nil, err
	}

	subs, err := s.rps.GetMySubs(ctx, id)
	if err != nil {
		return nil, nil, nil, err
	}

	return user, posts, subs, nil
}

func (s *Service) DeleteUser(ctx context.Context, id string) error {
	return s.rps.Delete(ctx, id)
}

func (s *Service) UpdateUser(ctx context.Context, id string, user *models.UserUpdate) error {
	return s.rps.Update(ctx, id, user)
}

func (s *Service) GetNewUsers(ctx context.Context) ([]*models.LastUsers, error) {
	return s.rps.GetLastUsersIDs(ctx)
}

/*

func (s *Service) AddSubscription(c echo.Context) error {
	userFromJwt := c.Get("user").(*jwt.Token) //why c.Get("user") to get auth header
	claims := userFromJwt.Claims.(jwt.MapClaims)
	idFromParam := claims["id"].(string)

	subid := c.Param("id")

	err := s.rps.AddSubscriprion(c.Request().Context(), subid, idFromParam)
	if err != nil {
		return echo.NewHTTPError(405, err.Error())
	}
	return c.JSON(200, http.NoBody)

}

func (s *Service) DeleteSubscription(c echo.Context) error {
	userFromJwt := c.Get("user").(*jwt.Token) //why c.Get("user") to get auth header
	claims := userFromJwt.Claims.(jwt.MapClaims)
	idFromParam := claims["id"].(string)

	subid := c.Param("id")

	err := s.rps.DeleteSubscription(c.Request().Context(), subid, idFromParam)
	if err != nil {
		return echo.NewHTTPError(405, err.Error())
	}
	return c.JSON(200, http.NoBody)

}

func (s *Service) GetLastUsers(c echo.Context) error {
	user, err := s.rps.GetLastUsers(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

// UpdateUser update user in db
func (s *Service) UpdateUser(c echo.Context) error {

	userFromJwt := c.Get("user").(*jwt.Token) //why c.Get("user") to get auth header
	claims := userFromJwt.Claims.(jwt.MapClaims)
	idFromJwt := claims["id"].(string)

	newClaims := model.UserUpdate{}
	err := json.NewDecoder(c.Request().Body).Decode(&newClaims)
	if err != nil {
		log.Errorf("failed parse json, %w", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	err = s.rps.UpdateUser(c.Request().Context(), idFromJwt, &newClaims)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, nil)
}

func (s *Service) GetReviewFeed(c echo.Context) error {
	return nil
}

func (s *Service) MySubscriptions(c echo.Context) error {
	userFromJwt := c.Get("user").(*jwt.Token) //why c.Get("user") to get auth header
	claims := userFromJwt.Claims.(jwt.MapClaims)
	idFromJwt := claims["id"].(string)

	users, err := s.rps.GetSubs(c.Request().Context(), idFromJwt)
	if err != nil {
		return echo.NewHTTPError(404, err.Error())
	}

	return c.JSON(200, users)
}

// DeleteUser delete user from db
func (s *Service) DeleteUser(c echo.Context) error {
	userFromJwt := c.Get("user").(*jwt.Token) //why c.Get("user") to get auth header
	claims := userFromJwt.Claims.(jwt.MapClaims)
	idFromJwt := claims["id"].(string)

	err := s.rps.DeleteUser(c.Request().Context(), idFromJwt)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	c.SetCookie(&http.Cookie{
		Name:   "token",
		Path:   s.rps.Co.CookiePath,
		Value:  "",
		MaxAge: -1,
	})

	return c.JSON(http.StatusOK, nil)
}*/
