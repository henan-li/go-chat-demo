package model

import "errors"

var (
	ERROR_USER_NOTEXISTS = errors.New("NO THIS USER")
	ERROR_USER_EXISTS = errors.New("USER REGISTERED")
	ERROR_USER_PWD = errors.New("WRONG INPUT")
)
