package model

import "errors"

var (
	ErrorNoPosts          = errors.New("you dont have any posts")
	ErrorUserDoesntExist  = errors.New("user with this email/username doesnt exist")
	ErrorUserAlreadyExist = errors.New("user with this email/username already exist")
)
