package models

import "errors"

var (
	ErrorNoPosts          = errors.New(" У вас нету ни одной рецензии")
	ErrorUserDoesntExist  = errors.New(" Пользователь не зарегестрирован")
	ErrorUserAlreadyExist = errors.New(" Этот никнейм/email уже занят")
	ErrorPasswordIsIncorrect = errors.New(" Неверный пароль")
)
