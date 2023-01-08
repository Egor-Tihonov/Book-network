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

// MyCookie ...
type MyCookie struct {
	CookieName   string
	CookieMaxAge int
	CookiePath   string
}

//User for response to ckient
type User struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Status   string `json:"status"`
}

type UserUpdate struct {
	Status  string `json:"status"`
	City    string `json:"city"`
	Country string `json:"country"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
	Bthsday string `json:"bthsday"`
}
