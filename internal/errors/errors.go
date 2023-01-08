/*creatte user errors models(add user errors that already exist)*/
package errmodel

import "errors"

var (
	ErrorNoPosts = errors.New("you dont have any posts")
	ErrorUserDoesntExist  = errors.New("user with this id/username doesnt exist")
	ErrorUserAlreadyExist = errors.New("user with this username already exist")
)
