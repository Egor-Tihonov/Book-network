package model

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type UserModel struct {
	Id           string `json:"id"`
	Username     string `json:"username"`
	Name         string `json:"name"`
	Password     string `bson,json:"password"`
	RefreshToken string `bson,json:"refreshToken"`
}

type PostModel struct {
	UserId     string `json:"user-id"`
	NameOfBook string `json:"name-of-book"`
	Content    string `json:"content"`
}

type AuthentcationForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//select * from postmodel
type Ctxkey string

type JWTClaims struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

var (
	//expirationTime work time
	ExpirationTime = time.Now().Add(1 * time.Minute)
)
