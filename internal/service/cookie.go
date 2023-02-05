package service

import (
	"context"
	"net/http"
	"time"
)

func (s *Service) setTokenCookie(name, token string, expiration time.Time, c context.Context) *http.Cookie {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = token
	cookie.Expires = expiration
	cookie.Path = "/"
	cookie.HttpOnly = true
	return cookie
}
