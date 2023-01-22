package service

import (
	"context"
	"net/http"
	"time"

	"github.com/Egor-Tihonov/Book-network/internal/model"
	"github.com/golang-jwt/jwt"
)


func (s *Service) GenerateTokensAndSetCookies(user *model.AuthUserModel, c context.Context) (cookieToken , cookieUser *http.Cookie, err error) {
	accessToken, exp, err := s.generateAccessToken(user)
	if err != nil {
		return
	}

	cookieToken = s.setTokenCookie(s.Co.CookieTokenName, accessToken, exp, c)
	cookieUser = s.setUserCookie(user, exp, c)
	return
}

func (s *Service) generateAccessToken(user *model.AuthUserModel) (string, time.Time, error) {
	expirationTime := time.Now().Add(1 * time.Hour)

	return s.generateToken(user, expirationTime, []byte(s.JWTKey))
}

func (s *Service) generateToken(user *model.AuthUserModel, expirationTime time.Time, secret []byte) (string, time.Time, error) {
	// Create the JWT claims, which includes the username and expiry time
	claims := &model.JWTClaims{
		ID:    user.ID,
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create the JWT string
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", time.Now(), err
	}

	return tokenString, expirationTime, nil
}
