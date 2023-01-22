// Package Service ...
package service

import (
	"context"
	"errors"
	"net/mail"
	"time"

	"github.com/Egor-Tihonov/Book-network/internal/model"
	"golang.org/x/crypto/bcrypt"
)

// RegistrationUser register new user, hash his password
func (s *Service) RegistrationUser(ctx context.Context, user *model.UserModel) error {
	hashPassword, err := s.hashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashPassword
	return s.rps.Create(ctx, user, time.Now())
}

// Authentication check user password, extradition jwt tokens
func (s *Service) Authentication(ctx context.Context, authForm *model.AuthenticationForm) (*model.AuthUserModel, error) {
	_, err := mail.ParseAddress(authForm.AuthString)
	var user *model.AuthUserModel
	if err != nil {
		user, err = s.rps.GetAuthByUsername(ctx, authForm.AuthString)
		if err != nil {
			return nil, err
		}
	} else {
		user, err = s.rps.GetAuthByEmail(ctx, authForm.AuthString)
		if err != nil {
			return nil, err
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(authForm.Password))
	if err != nil {
		err = errors.New("password is incorrect")
		return nil, err
	}

	return user, nil
}

// hashPassword ...
func (s *Service) hashPassword(password string) (newHashPassword string, err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return
	}
	newHashPassword = string(bytes)
	return
}
