// Package model ...
package model

import (
	"github.com/golang-jwt/jwt"
)

// UserModel ...
type UserModel struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Password string `bson,json:"password"`
}

// AuthenticationForm ...
type AuthenticationForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// JWTClaims json web tokens claims
type JWTClaims struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

type MyCookie struct {
	CookieName   string
	CookieMaxAge int
	CookiePath   string
}
