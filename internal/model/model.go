package model

import (
	"time"

	"github.com/golang-jwt/jwt"
)

//UserModel ...
type UserModel struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Password string `bson,json:"password"`
}

//AuthenticationForm
type AuthenticationForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//JWTClaims json web tokens claims
type JWTClaims struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

//expirationTime work time
var ExpirationTime = time.Now().Add(1 * time.Minute)
