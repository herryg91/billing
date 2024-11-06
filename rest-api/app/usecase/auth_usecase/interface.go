package auth_usecase

import (
	"errors"

	"github.com/herryg91/billing/rest-api/app/entity"
)

var ErrUnexpected = errors.New("Unexpected internal error")
var ErrLoginEmailNotRegistered = errors.New("Email hasn't registered")
var ErrLoginInvalidPassword = errors.New("Invalid password")
var ErrLoginFailedGenerateToken = errors.New("Failed to generate token")
var ErrUserNotFound = errors.New("user not found")

type UseCase interface {
	Login(email, password string) (token, refresh_token string, err error)
	GetAuthenticatedUser(email string) (*entity.User, error)

	Register(email, password, name string) error
}
