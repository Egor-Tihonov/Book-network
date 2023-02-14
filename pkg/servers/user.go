package servers

import (
	"context"
	"net/http"

	"github.com/Egor-Tihonov/Book-network/pkg/models"
	"github.com/Egor-Tihonov/Book-network/pkg/pb"
)

// GetUser get info about user from db
func (s *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	if err := s.Rps.Create(ctx, &models.UserModel{
		ID:       req.Id,
		Name:     req.Name,
		Email:    req.Email,
		Username: req.Username,
	}); err != nil {
		return &pb.CreateUserResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, err
	}
	return &pb.CreateUserResponse{
		Status: http.StatusOK,
	}, nil
} /*

func (s *Server) GetUser(c echo.Context) error {
	userFromJwt := c.Get("user").(*jwt.Token) //why c.Get("user") to get auth header
	claims := userFromJwt.Claims.(jwt.MapClaims)
	idFromParam := claims["id"].(string)

	user, err := s.Rps.GetUser(c.Request().Context(), idFromParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

func (s *Server) GetOtherUser(c echo.Context) error {
	id := c.Param("id")

	userFromJwt := c.Get("user").(*jwt.Token) //why c.Get("user") to get auth header
	claims := userFromJwt.Claims.(jwt.MapClaims)
	idFromParam := claims["id"].(string)

	user, err := s.Rps.GetUser(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	exist, err := s.Rps.CheckSubs(c.Request().Context(), id, idFromParam)
	if err != nil {
		return echo.NewHTTPError(404, err.Error())
	}

	response := &model.GetOtherUserResponse{
		User:         user,
		Subscription: exist,
	}
	return c.JSON(http.StatusOK, response)
}

func (s *Server) AddSubscription(c echo.Context) error {
	userFromJwt := c.Get("user").(*jwt.Token) //why c.Get("user") to get auth header
	claims := userFromJwt.Claims.(jwt.MapClaims)
	idFromParam := claims["id"].(string)

	subid := c.Param("id")

	err := s.Rps.AddSubscriprion(c.Request().Context(), subid, idFromParam)
	if err != nil {
		return echo.NewHTTPError(405, err.Error())
	}
	return c.JSON(200, http.NoBody)

}

func (s *Server) DeleteSubscription(c echo.Context) error {
	userFromJwt := c.Get("user").(*jwt.Token) //why c.Get("user") to get auth header
	claims := userFromJwt.Claims.(jwt.MapClaims)
	idFromParam := claims["id"].(string)

	subid := c.Param("id")

	err := s.Rps.DeleteSubscription(c.Request().Context(), subid, idFromParam)
	if err != nil {
		return echo.NewHTTPError(405, err.Error())
	}
	return c.JSON(200, http.NoBody)

}

func (s *Server) GetLastUsers(c echo.Context) error {
	user, err := s.Rps.GetLastUsers(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

// UpdateUser update user in db
func (s *Server) UpdateUser(c echo.Context) error {

	userFromJwt := c.Get("user").(*jwt.Token) //why c.Get("user") to get auth header
	claims := userFromJwt.Claims.(jwt.MapClaims)
	idFromJwt := claims["id"].(string)

	newClaims := model.UserUpdate{}
	err := json.NewDecoder(c.Request().Body).Decode(&newClaims)
	if err != nil {
		log.Errorf("failed parse json, %w", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	err = s.Rps.UpdateUser(c.Request().Context(), idFromJwt, &newClaims)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, nil)
}

func (s *Server) GetReviewFeed(c echo.Context) error {
	return nil
}

func (s *Server) MySubscriptions(c echo.Context) error {
	userFromJwt := c.Get("user").(*jwt.Token) //why c.Get("user") to get auth header
	claims := userFromJwt.Claims.(jwt.MapClaims)
	idFromJwt := claims["id"].(string)

	users, err := s.Rps.GetSubs(c.Request().Context(), idFromJwt)
	if err != nil {
		return echo.NewHTTPError(404, err.Error())
	}

	return c.JSON(200, users)
}

// DeleteUser delete user from db
func (s *Server) DeleteUser(c echo.Context) error {
	userFromJwt := c.Get("user").(*jwt.Token) //why c.Get("user") to get auth header
	claims := userFromJwt.Claims.(jwt.MapClaims)
	idFromJwt := claims["id"].(string)

	err := s.Rps.DeleteUser(c.Request().Context(), idFromJwt)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	c.SetCookie(&http.Cookie{
		Name:   "token",
		Path:   s.Rps.Co.CookiePath,
		Value:  "",
		MaxAge: -1,
	})

	return c.JSON(http.StatusOK, nil)
}*/
