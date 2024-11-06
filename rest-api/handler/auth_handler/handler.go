package auth_handler

import (
	"context"
	"errors"
	"net/http"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/herryg91/billing/rest-api/app/usecase/auth_usecase"
	pbAuth "github.com/herryg91/billing/rest-api/handler/grst/auth"
	"github.com/herryg91/billing/rest-api/pkg/interceptor/usertoken_interceptor"
	grst_errors "github.com/herryg91/cdd/grst/errors"
	"google.golang.org/grpc/codes"
)

type handler struct {
	pbAuth.UnimplementedUserAuthApiServer
	auth_uc auth_usecase.UseCase
}

func NewAuthHandler(auth_uc auth_usecase.UseCase) pbAuth.UserAuthApiServer {
	return &handler{auth_uc: auth_uc}
}

func (h *handler) Login(ctx context.Context, req *pbAuth.LoginReq) (*pbAuth.UserToken, error) {
	if err := pbAuth.ValidateRequest(req); err != nil {
		return nil, err
	}

	token, refresh_token, err := h.auth_uc.Login(req.Email, req.Password)
	if err != nil {
		if errors.Is(err, auth_usecase.ErrLoginEmailNotRegistered) {
			return nil, grst_errors.New(http.StatusBadRequest, codes.InvalidArgument, 60102, err.Error(), &grst_errors.ErrorDetail{Code: 1, Field: "email", Message: "Email `" + req.Email + "` not found"})
		} else if errors.Is(err, auth_usecase.ErrLoginInvalidPassword) {
			return nil, grst_errors.New(http.StatusBadRequest, codes.InvalidArgument, 60103, err.Error(), &grst_errors.ErrorDetail{Code: 2, Field: "password", Message: "You've entered invalid password"})
		}
		return nil, grst_errors.New(http.StatusInternalServerError, codes.Internal, 60104, err.Error())
	}

	return &pbAuth.UserToken{
		AuthToken:    token,
		RefreshToken: refresh_token,
	}, nil
}

func (h *handler) GetAuthenticated(ctx context.Context, req *empty.Empty) (*pbAuth.AuthenticatedUser, error) {
	user_ctx, err := usertoken_interceptor.GetUserContext(ctx)
	if err != nil {
		return nil, err
	}

	u, err := h.auth_uc.GetAuthenticatedUser(user_ctx.Email)
	if err != nil {
		if errors.Is(err, auth_usecase.ErrUserNotFound) {
			return nil, grst_errors.New(http.StatusNotFound, codes.NotFound, 404, err.Error())
		}
		return nil, grst_errors.New(http.StatusInternalServerError, codes.Internal, 500, err.Error())
	}

	return &pbAuth.AuthenticatedUser{
		Id:    int32(u.Id),
		Email: u.Email,
		Name:  u.Name,
	}, nil
}
