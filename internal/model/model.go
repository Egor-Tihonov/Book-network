// Package model ...
package model

import (
	"github.com/golang-jwt/jwt"
)

// UserModel for create user
type UserModel struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Password string `bson,json:"password"`
	Email    string `json:"email"`
}

type AuthUserModel struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Password string `bson,json:"password"`
}

// AuthenticationForm ...
type AuthenticationForm struct {
	AuthString string `json:"authString"`
	Password string `json:"password"`
}

// JWTClaims json web tokens claims
type JWTClaims struct {
	ID    string `json:"id"`
	Email string `json:"email"`
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
	Email    string `json:"email"`
}

type UserUpdate struct {
	Status  string `json:"status"`
	City    string `json:"city"`
	Country string `json:"country"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
	Bthsday string `json:"bthsday"`
}
