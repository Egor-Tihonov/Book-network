package services

/*
func (s *Service) CreatePost(ctx context.Context, id string) error {
	return s.rps.CreatePost(ctx, id)
}

/*
func (s *Service) GetPosts(c echo.Context) error {
	id := c.Param("id")
	posts, err := s.se.GetPosts(c.Request().Context(), id)
	if err != nil {
		logrus.Errorf("get all user posts error: %w", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, posts)
}

func (s *Service) GetMyPosts(c echo.Context) error {
	userFromJwt := c.Get("user").(*jwt.Token) //why c.Get("user") to get auth header
	claims := userFromJwt.Claims.(jwt.MapClaims)
	idFromJwt := claims["id"].(string)

	posts, err := s.se.GetPosts(c.Request().Context(), idFromJwt)
	if err != nil {
		logrus.Errorf("get my posts error: %w", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, posts)
}

func (s *Service) GetPost(c echo.Context) error {
	userFromJwt := c.Get("user").(*jwt.Token) //why c.Get("user") to get auth header
	claims := userFromJwt.Claims.(jwt.MapClaims)
	idFromJwt := claims["id"].(string)

	postID := c.Param("id")

	post, err := s.se.GetPost(c.Request().Context(), idFromJwt, postID)
	if err != nil {
		if errors.Is(err, models.ErrorNoPosts) {
			return echo.NewHTTPError(404, err.Error())
		}
		logrus.Errorf("get one post error: %w", err)
		return echo.NewHTTPError(405, err.Error())
	}

	return c.JSON(http.StatusOK, post)
}

func (s *Service) GetLastPosts(c echo.Context) error {
	userFromJwt := c.Get("user").(*jwt.Token) //why c.Get("user") to get auth header
	claims := userFromJwt.Claims.(jwt.MapClaims)
	idFromJwt := claims["id"].(string)

	lastPosts, err := s.se.GetLast(c.Request().Context(), idFromJwt)
	if err != nil {
		if errors.Is(err, models.ErrorNoPosts) {
			return echo.NewHTTPError(404, err.Error())
		}
		logrus.Errorf("get last post error: %w", err)
		return echo.NewHTTPError(405, err.Error())
	}

	return c.JSON(200, lascontext
}

func (s *Service) GetAllPosts(c echo.Context) error {
	userFromJwt := c.Get("user").(*jwt.Token) //why c.Get("user") to get auth header
	claims := userFromJwt.Claims.(jwt.MapClaims)
	idFromJwt := claims["id"].(string)

	posts, err := s.se.GetAllPosts(c.Request().Context(), idFromJwt)
	if err != nil {
		logrus.Errorf("get my feed posts error: %w", err)
		return echo.NewHTTPError(404, err.Error())
	}
	return c.JSON(200, posts)
}
*/
