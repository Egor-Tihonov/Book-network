// Package server ...
package server

import (
	"context"
	"errors"
	"time"

	"github.com/Egor-Tihonov/Book-network/internal/model"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var (
	// ErrorEmptyUsername empty username
	ErrorEmptyUsername = errors.New("username couldnt be empty")
	// ErrorComparePassword false password
	ErrorComparePassword = errors.New("passwrod not correct")
	// JwtKey secure key
	// JwtKey = []byte("super-key")
)

// RegistrationUser register new user, hash his password
func (s *Server) RegistrationUser(ctx context.Context, person *model.UserModel) error {
	err := hashPassword(person)
	if err != nil {
		return err
	}
	if person.Username == "" {
		return ErrorEmptyUsername
	}
	return s.rps.Create(ctx, person)
}

// Authentication check user password, extradition jwt tokens
func (s *Server) Authentication(ctx context.Context, authForm *model.AuthenticationForm) (token string, err error) {
	user, err := s.rps.GetAuth(ctx, authForm.Username)
	if err != nil {
		return
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
func (s *Server) generateJWT(user *model.UserModel) (accessTokenStr string, err error) {
	claims := model.JWTClaims{
		ID:       user.ID,
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(1 * time.Minute).Unix(),
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
func hashPassword(person *model.UserModel) (err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(person.Password), 14)
	if err != nil {
		return
	}
	person.Password = string(bytes)
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
