package usertoken_handler

import (
	"context"
	"errors"
	"net/http"

	"github.com/herryg91/billing/rest-api/app/entity"
	"github.com/herryg91/billing/rest-api/app/usecase/usertoken_usecase"
	pbUserToken "github.com/herryg91/billing/rest-api/handler/grst/usertoken"
	grst_errors "github.com/herryg91/cdd/grst/errors"
	"google.golang.org/grpc/codes"
)

type handler struct {
	pbUserToken.UnimplementedUsertokenAPIServer
	usertoken_uc usertoken_usecase.UseCase[entity.UserTokenClaim]
}

func New(usertoken_uc usertoken_usecase.UseCase[entity.UserTokenClaim]) pbUserToken.UsertokenAPIServer {
	return &handler{usertoken_uc: usertoken_uc}
}

// GenerateAuthToken. Code 10000
func (h *handler) GenerateToken(ctx context.Context, req *pbUserToken.GenerateTokenReq) (*pbUserToken.UserToken, error) {
	if err := pbUserToken.ValidateRequest(req); err != nil {
		return nil, err
	}

	authToken, refreshToken, err := h.usertoken_uc.GenerateToken(entity.UserTokenClaim{
		UserId: int(req.UserId),
		Email:  req.Email,
	})
	if err != nil {
		return nil, grst_errors.New(http.StatusInternalServerError, codes.Internal, 10001, err.Error())
	}

	return &pbUserToken.UserToken{
		AuthToken:    authToken,
		RefreshToken: refreshToken,
	}, nil
}

// RefreshToken. Code 11000
func (h *handler) RefreshToken(ctx context.Context, req *pbUserToken.RefreshTokenReq) (*pbUserToken.UserToken, error) {
	if err := pbUserToken.ValidateRequest(req); err != nil {
		return nil, err
	}

	oldTokenClaim, err := h.usertoken_uc.ValidateRefreshToken(req.RefreshToken)
	if err != nil {
		return nil, grst_errors.New(http.StatusForbidden, codes.PermissionDenied, 11001, err.Error())
	}

	tokenClaim := entity.UserTokenClaim{UserId: oldTokenClaim.UserId, Email: oldTokenClaim.Email}
	authToken, refreshToken, err := h.usertoken_uc.GenerateToken(tokenClaim)
	if err != nil {
		return nil, grst_errors.New(http.StatusInternalServerError, codes.Internal, 11002, err.Error())
	}

	return &pbUserToken.UserToken{
		AuthToken:    authToken,
		RefreshToken: refreshToken,
		TokenType:    "bearer",
	}, nil
}

// ValidateToken. Code 12000
func (h *handler) ValidateToken(ctx context.Context, req *pbUserToken.ValidateTokenReq) (*pbUserToken.UserTokenClaim, error) {
	if err := pbUserToken.ValidateRequest(req); err != nil {
		return nil, err
	}

	claim, err := h.usertoken_uc.ValidateToken(req.AuthToken)
	if err != nil {
		if errors.Is(err, usertoken_usecase.ErrTokenExpired) {
			return nil, grst_errors.New(http.StatusUnauthorized, codes.Unauthenticated, 12001, err.Error())
		}
		return nil, grst_errors.New(http.StatusInternalServerError, codes.Internal, 12002, err.Error())
	}

	return &pbUserToken.UserTokenClaim{
		UserId: int32(claim.UserId),
		Email:  claim.Email,
	}, nil

}
