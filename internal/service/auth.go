// Package Service ...
package service

import (
	"context"
	"errors"
	"net/mail"
	"time"

	"github.com/Egor-Tihonov/Book-network/internal/model"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// RegistrationUser register new user, hash his password
func (s *Service) RegistrationUser(ctx context.Context, user *model.UserModel) error {
	err := s.hashPassword(user)
	if err != nil {
		return err
	}
	return s.rps.Create(ctx, user)
}

// Authentication check user password, extradition jwt tokens
func (s *Service) Authentication(ctx context.Context, authForm *model.AuthenticationForm) (token string, err error) {
	_, err = mail.ParseAddress(authForm.AuthString)
	var user *model.AuthUserModel
	if err != nil {
		user, err = s.rps.GetAuthByUsername(ctx, authForm.AuthString)
		if err != nil {
			return "", err
		}
	} else {
		user, err = s.rps.GetAuthByEmail(ctx, authForm.AuthString)
		if err != nil {
			return "", err
		}
	}
	err = comparePassword(authForm.Password, user.Password)
	if err != nil {
		return
	}
	token, err = s.generateJWT(user)
	if err != nil {
		return
	}
	return
}

// generateJWT ...
func (s *Service) generateJWT(user *model.AuthUserModel) (accessTokenStr string, err error) {
	claims := model.JWTClaims{
		ID:    user.ID,
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessTokenStr, err = token.SignedString(s.JWTKey)
	if err != nil {
		return
	}
	return
}

// hashPassword ...
func (s *Service) hashPassword(user *model.UserModel) (err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return
	}
	user.Password = string(bytes)
	return
}

// comparePasswrod ...
func comparePassword(password, hashedPassword string) (err error) {
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		err = errors.New("password is incorrect")
		return
	}
	return
}
