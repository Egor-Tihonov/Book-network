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
})

func JWTErrorChecker(err error, c echo.Context) error {
	// Redirects to the signIn form.
	c.SetCookie(&http.Cookie{
		Name:   "user",
		Path:   "/",
		Value:  "",
		MaxAge: -1,
	})

	c.SetCookie(&http.Cookie{
		Name:   "token",
		Path:   "/",
		Value:  "",
		MaxAge: -1,
	})
	return echo.NewHTTPError(401)
}
