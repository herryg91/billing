package usertoken_interceptor

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/herryg91/billing/rest-api/app/entity"
	usertoken_usecase "github.com/herryg91/billing/rest-api/app/usecase/usertoken_usecase"
	grst_errors "github.com/herryg91/cdd/grst/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

type AuthConditionRole string

func NewAuthConditionRole(roles []string, fallback AuthConditionRole) AuthConditionRole {
	for _, role := range roles {
		if role == "*" {
			return AuthConditionRole_All
		}
	}
	return fallback
}

const AuthConditionRole_All AuthConditionRole = "*"

type AuthCondition struct {
	NeedAuth bool
	Role     AuthConditionRole
}

func UnaryServerInterceptor(usertoken_uc usertoken_usecase.UseCase[entity.UserTokenClaim], auth_conditions map[string]AuthCondition) grpc.UnaryServerInterceptor {

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		if _, ok := auth_conditions[info.FullMethod]; !ok {
			return handler(ctx, req)
		}
		auth_condition := auth_conditions[info.FullMethod]
		if auth_condition.NeedAuth == false {
			return handler(ctx, req)
		}

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, grst_errors.New(http.StatusInternalServerError, codes.Internal, 98001, "Failed to parse context")
		} else if len(md["authorization"]) <= 0 {
			return nil, grst_errors.New(http.StatusBadRequest, codes.InvalidArgument, 98002, "Authorization on header is required")
		}

		token := md["authorization"][0]
		if len(token) < 7 {
			return nil, grst_errors.New(http.StatusBadRequest, codes.InvalidArgument, 98003, "Invalid token format")
		} else if strings.ToLower(token[:7]) != "bearer " {
			return nil, grst_errors.New(http.StatusBadRequest, codes.InvalidArgument, 98004, "Invalid token type. Need bearer token type")
		}
		token = token[7:]

		successResp, err := usertoken_uc.ValidateToken(token)
		if err != nil {
			if errors.Is(err, usertoken_usecase.ErrTokenExpired) {
				return nil, grst_errors.New(http.StatusUnauthorized, codes.Unauthenticated, 401, err.Error())
			}
			return nil, grst_errors.New(http.StatusInternalServerError, codes.Internal, 500, err.Error())
		}

		// Check Role

		// if auth_condition.Role == AuthConditionRole_Superadmin {
		// 	is_admin := admin_emails[successResp.Email]
		// 	if !is_admin {
		// 		return nil, grst_errors.New(http.StatusForbidden, codes.PermissionDenied, 403, "Role mismatch")
		// 	}
		// }

		ctx = SetUserContext(ctx, UserContext{
			UserId: successResp.UserId,
			Email:  successResp.Email,
		})
		return handler(ctx, req)
	}
}
