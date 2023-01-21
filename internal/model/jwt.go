package model

import "github.com/golang-jwt/jwt"

// JWTClaims json web tokens claims
type JWTClaims struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}
