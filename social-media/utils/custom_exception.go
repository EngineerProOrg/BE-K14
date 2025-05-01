package utils

import "errors"

var ErrInvalidLogin = errors.New("email or password is incorrect")
var ErrUserDoesNotExist = errors.New("user does not exist")
