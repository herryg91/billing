package usertoken_interceptor

import (
	"context"
	"fmt"
	"strconv"

	"google.golang.org/grpc/metadata"
)

type ContextKey string

func (key ContextKey) String() string {
	return string(key)
}

const (
	CONTEXT_USER_ID    ContextKey = "usertoken.user.id"
	CONTEXT_USER_EMAIL ContextKey = "usertoken.user.email"
)

type UserContext struct {
	UserId int
	Email  string
}

func GetUserContext(ctx context.Context) (UserContext, error) {
	result := UserContext{}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return result, fmt.Errorf("context is invalid: %s", fmt.Sprint(ctx))
	}
	result.UserId = parseMetadataToInt(md, CONTEXT_USER_ID)
	result.Email = parseMetadataToString(md, CONTEXT_USER_EMAIL)
	return result, nil
}

func SetUserContext(ctx context.Context, userCtx UserContext) context.Context {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		md = metadata.New(map[string]string{})
	}

	md.Set(CONTEXT_USER_ID.String(), strconv.Itoa(userCtx.UserId))
	md.Set(CONTEXT_USER_EMAIL.String(), userCtx.Email)

	ctx = metadata.NewIncomingContext(ctx, md)

	return ctx
}

func parseMetadataToString(md metadata.MD, key ContextKey) string {
	result := md.Get(key.String())
	if len(result) > 0 {
		return result[0]
	}
	return ""

}
func parseMetadataToInt(md metadata.MD, key ContextKey) int {
	result := md.Get(key.String())
	if len(result) > 0 {
		resultInt, _ := strconv.Atoi(result[0])
		return resultInt
	}
	return -1

}
