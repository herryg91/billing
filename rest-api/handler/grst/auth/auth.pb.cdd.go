// Code generated by protoc-gen-cdd. DO NOT EDIT.
// source: auth.proto
package auth

import (
	"net/http"
	"strings"

	"github.com/herryg91/cdd/grst"
	grst_errors "github.com/herryg91/cdd/grst/errors"
	"google.golang.org/grpc"

	"github.com/mcuadros/go-defaults"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"gopkg.in/validator.v2"
)

type fullMethods struct {
	UserAuthApi_Login            string
	UserAuthApi_GetAuthenticated string
}

var FullMethods = fullMethods{
	UserAuthApi_Login:            "/auth.UserAuthApi/Login",
	UserAuthApi_GetAuthenticated: "/auth.UserAuthApi/GetAuthenticated",
}

var NeedAuthFullMethods = []string{

	"/auth.UserAuthApi/GetAuthenticated",
}

type AuthConfig struct {
	NeedAuth bool
	Roles    []string
}

var AuthConfigFullMethods = map[string]AuthConfig{
	"/auth.UserAuthApi/Login":            AuthConfig{NeedAuth: false, Roles: []string{"*"}},
	"/auth.UserAuthApi/GetAuthenticated": AuthConfig{NeedAuth: true, Roles: []string{"*"}},
}

var NeedApiKeyFullMethods = []string{}

func ValidateRequest(req interface{}) error {
	defaults.SetDefaults(req)
	if errs := validator.Validate(req); errs != nil {
		validateError := []*grst_errors.ErrorDetail{}
		for field, err := range errs.(validator.ErrorMap) {
			errMessage := strings.Replace(err.Error(), "{field}", field, -1)
			validateError = append(validateError, &grst_errors.ErrorDetail{Code: 999, Field: field, Message: errMessage})
		}
		return grst_errors.New(http.StatusBadRequest, codes.InvalidArgument, 999, "Validation Error", validateError...)
	}

	return nil
}

/*==================== UserAuthApi Section ====================*/

func RegisterUserAuthApiGrstServer(grpcRestServer *grst.Server, hndl UserAuthApiServer) {

	forward_UserAuthApi_Login_0 = grpcRestServer.GetForwardResponseMessage()

	forward_UserAuthApi_GetAuthenticated_0 = grpcRestServer.GetForwardResponseMessage()

	RegisterUserAuthApiServer(grpcRestServer.GetGrpcServer(), hndl)
	grpcRestServer.RegisterRestHandler(RegisterUserAuthApiHandler)
}

func NewUserAuthApiGrstClient(serverHost string, creds *credentials.TransportCredentials, dialOpts ...grpc.DialOption) (UserAuthApiClient, error) {
	opts := []grpc.DialOption{}
	opts = append(opts, grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(1024*1024*20)))
	opts = append(opts, grpc.WithMaxMsgSize(1024*1024*20))
	if creds == nil {
		opts = append(opts, grpc.WithInsecure())
	} else {
		opts = append(opts, grpc.WithTransportCredentials(*creds))
	}
	opts = append(opts, dialOpts...)
	grpcConn, err := grpc.Dial(serverHost, opts...)
	return NewUserAuthApiClient(grpcConn), err
}
