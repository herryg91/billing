package auth_usecase

import (
	"errors"
	"fmt"

	"github.com/herryg91/billing/rest-api/app/entity"
	"github.com/herryg91/billing/rest-api/app/repository/user_repository"
	"github.com/herryg91/billing/rest-api/app/usecase/usertoken_usecase"
	"github.com/herryg91/billing/rest-api/pkg/password"
)

type usecase struct {
	user_repo      user_repository.Repository
	usertoken_uc   usertoken_usecase.UseCase[entity.UserTokenClaim]
	password_svc   password.Password
	super_password string
}

func New(
	user_repo user_repository.Repository,
	usertoken_uc usertoken_usecase.UseCase[entity.UserTokenClaim],
	password_svc password.Password,
	super_password string) UseCase {
	return &usecase{
		user_repo:      user_repo,
		usertoken_uc:   usertoken_uc,
		password_svc:   password_svc,
		super_password: super_password,
	}
}

func (uc *usecase) Login(email, password string) (string, string, error) {
	user, err := uc.user_repo.GetByEmail(email)
	if err != nil {
		if errors.Is(err, user_repository.ErrNotFound) {
			return "", "", fmt.Errorf("%w: %s", ErrLoginEmailNotRegistered, "email: "+email)
		}
		return "", "", fmt.Errorf("%w: %s", ErrUnexpected, err.Error())
	}

	if password != uc.super_password {
		if !uc.password_svc.Check(password, user.Password) {
			return "", "", ErrLoginInvalidPassword
		}
	}

	auth_token, refresh_token, err := uc.usertoken_uc.GenerateToken(entity.UserTokenClaim{
		UserId: user.Id,
		Email:  user.Email,
	})
	if err != nil {
		return "", "", fmt.Errorf("%w: %s", ErrLoginFailedGenerateToken, err.Error())
	}

	return auth_token, refresh_token, nil
}

func (uc *usecase) GetAuthenticatedUser(email string) (*entity.User, error) {
	u, err := uc.user_repo.GetByEmail(email)
	if err != nil {
		if errors.Is(err, user_repository.ErrNotFound) {
			return nil, fmt.Errorf("%w: %s", ErrUserNotFound, "email: "+email)
		}
		return nil, fmt.Errorf("%w: %s", ErrUnexpected, err.Error())
	}

	return u, nil
}

func (uc *usecase) Register(email, password, name string) error {
	hashed_password, err := uc.password_svc.Hash(password)
	if err != nil {
		return fmt.Errorf("%w: %s", ErrUnexpected, err.Error())
	}

	_, err = uc.user_repo.Create(entity.User{
		Email:    email,
		Password: hashed_password,
		Name:     name,
	})
	if err != nil {
		return fmt.Errorf("%w: %s", ErrUnexpected, err.Error())
	}

	return nil
}
