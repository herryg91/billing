package user_repository

import "errors"

var ErrNotFound = errors.New("User not found")
var ErrDuplicateEmail = errors.New("Email is duplicate")
