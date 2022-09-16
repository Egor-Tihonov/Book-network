package server

import (
	"context"
	"errors"
	"net/http"

	"github.com/Egor-Tihonov/Book-network/internal/model"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

var (
	//ErrorEmptyUsername empty username
	ErrorEmptyUsername = errors.New("username couldnt be empty")
	//ErrorComparePassword false password
	ErrorComparePassword    = errors.New("passwrod not correct")
	ErrorStatusUnautharized = errors.New("Unauthorized")
	//JwtKey secure key
	JwtKey = []byte("super-key")
	tknStr string
)

//Registration register new user, hash his password
func (s *Server) RegistrationUser(ctx context.Context, person *model.UserModel) error {
	err := hashPassword(person)
	if err != nil {
		return err
	}
	if person.Username == "" {
		return ErrorEmptyUsername
	}
	return s.rps.CreateUser(ctx, person)
}

//Authentication check user password, extradition jwt tokens
func (s *Server) Authentcation(ctx context.Context, authForm *model.AuthentcationForm) (string, error) {
	user, err := s.rps.SelectUserAuth(ctx, authForm.Username)
	if err != nil {
		return "", err
	}
	err = comparePassword(authForm.Password, user.Password)
	if err != nil {
		return "", err
	}
	token, err := generateJWT(user)
	if err != nil {
		return "", err
	}
	return token, nil
}

//generateJWT ...
func generateJWT(user *model.UserModel) (string, error) {
	claims := &model.JWTClaims{
		Id:       user.Id,
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: model.ExpirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessTokenStr, err := token.SignedString(JwtKey)
	if err != nil {
		return "", err
	}
	return accessTokenStr, nil
}

//hashPassword ...
func hashPassword(person *model.UserModel) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(person.Password), 14)
	if err != nil {
		return err
	}
	person.Password = string(bytes)
	return nil
}

//comparePasswrod ...
func comparePassword(password, hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		err = errors.New("password is incorrect")
		return err
	}
	return nil
}

func (s *Server) Validation(c echo.Context) (model.JWTClaims, error) {
	cookie, err := c.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			return model.JWTClaims{}, ErrorStatusUnautharized
		}
		return model.JWTClaims{}, err
	}
	tknStr = cookie.Value
	claims := &model.JWTClaims{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return model.JWTClaims{}, err
		}
		return model.JWTClaims{}, err
	}
	if !tkn.Valid {
		return model.JWTClaims{}, err
	}
	return *claims, nil
}
