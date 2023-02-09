package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var IsLoggedIn = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey:              []byte("SUPER-KEY"),
	TokenLookup:             "cookie:token",
	ErrorHandlerWithContext: JWTErrorChecker,
	Skipper: func(c echo.Context) bool {
		switch c.Request().URL.Path {
		case "/login":
			return true
		case "/user/logout":
			return true
		default:
			return false
		}
	},
})

func JWTErrorChecker(err error, c echo.Context) error {
	c.SetCookie(&http.Cookie{
		Name:   "token",
		Path:   "/",
		Value:  "",
		MaxAge: -1,
	})
	return echo.NewHTTPError(401)
}
