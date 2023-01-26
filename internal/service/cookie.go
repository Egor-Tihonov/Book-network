package service

import (
	"context"
	"net/http"
	"time"

	"github.com/Egor-Tihonov/Book-network/internal/model"
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

func (s *Service) setUserCookie(user *model.AuthUserModel, expiration time.Time, c context.Context) *http.Cookie {
	cookie := new(http.Cookie)
	cookie.Name = s.Co.CookieUserName
	cookie.Value = user.ID
	cookie.Expires = expiration
	cookie.Path = "/"
	return cookie
}
