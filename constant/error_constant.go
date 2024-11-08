package constant

import "errors"

var (
	EMAIL_NOT_FOUND   = errors.New("email not found")
	EMAIL_IS_EMPTY    = errors.New("email is empty")
	PASSWORD_IS_EMPTY = errors.New("password is empty")
	PASSWORD_IS_WRONG = errors.New("password is wrong")
)
