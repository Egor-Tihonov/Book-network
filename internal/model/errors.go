package model

import "errors"

var (
	ErrorNoPosts          = errors.New("у вас нету ни одной рецензии")
	ErrorUserDoesntExist  = errors.New("пользователь с этим никнеймом/email не зарегестрирован")
	ErrorUserAlreadyExist = errors.New("этот никнейм/email уже занят")
)
