package domain

import "errors"

var (
	ErrorNoResult                     = errors.New("no result")
	ErrorUserWithEmailAlreadyExist    = errors.New("User with this email already exists")
	ErrorUserWithUsernameAlreadyExist = errors.New("User with this username already exists")
)
