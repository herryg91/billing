package usertoken_usecase

import (
	"errors"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/herryg91/billing/rest-api/pkg/helpers"
)

type UseCase[T any] interface {
	GenerateToken(claims T) (auth_token, refresh_token string, err error)
	ValidateToken(authToken string) (*T, error)
	ValidateRefreshToken(refreshToken string) (*T, error)
}

var ErrGenerateJwtToken = errors.New("failed generate token")
var ErrInvalidToken = errors.New("token is invalid")
var ErrParseMapClaims = errors.New("failed to parse Map Claims")
var ErrTokenExpired = errors.New("token has expired")

type usecase[T any] struct {
	authTokenSecret    string
	authTokenExpiry    int
	refreshTokenSecret string
	refreshTokenExpiry int
}

func New[T any](authTokenSecret string, authTokenExpiry int, refreshTokenSecret string, refreshTokenExpiry int) UseCase[T] {
	return &usecase[T]{
		authTokenSecret:    authTokenSecret,
		authTokenExpiry:    authTokenExpiry,
		refreshTokenSecret: refreshTokenSecret,
		refreshTokenExpiry: refreshTokenExpiry,
	}
}

func (uc *usecase[T]) GenerateToken(claims T) (auth_token, refresh_token string, err error) {
	auth_token, refresh_token, err = "", "", nil

	claims_map := helpers.StructToMap(claims)
	// Generate Refresh Token
	refresh_jwt_claims := RefreshTokenClaim{}.New(claims_map, uc.refreshTokenExpiry).ToJwtClaims()
	refresh_token, err = generateToken(refresh_jwt_claims, uc.refreshTokenSecret)
	if err != nil {
		return
	}

	// Generate Auth Token
	auth_jwt_claims := AuthTokenClaim{}.New(claims_map, refresh_token, uc.authTokenExpiry).ToJwtClaims()
	auth_token, err = generateToken(auth_jwt_claims, uc.authTokenSecret)
	return
}

func (uc *usecase[T]) ValidateToken(authToken string) (*T, error) {
	parsedToken, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(uc.authTokenSecret), nil
	})
	if err != nil {
		return nil, fmt.Errorf("<%w>: %s", ErrInvalidToken, err.Error())
	}

	atc := AuthTokenClaim{claims: map[string]interface{}{}}
	if jwtclaims, ok := parsedToken.Claims.(jwt.MapClaims); ok {
		atc = AuthTokenClaim{}.FromJwtClaims(jwtclaims)
	}
	if atc.IsExpired() {
		return nil, ErrTokenExpired
	}

	var out T
	helpers.MapToStruct(atc.claims, &out)

	return &out, nil
}

func (uc *usecase[T]) ValidateRefreshToken(refreshToken string) (*T, error) {
	parsedToken, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(uc.refreshTokenSecret), nil
	})
	if err != nil {
		return nil, fmt.Errorf("<%w>: %s", ErrInvalidToken, err.Error())
	}

	rtc := RefreshTokenClaim{}
	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok {
		rtc = RefreshTokenClaim{}.FromJwtClaims(claims)
	}

	if rtc.IsExpired() {
		return nil, ErrTokenExpired
	}

	var out T
	helpers.MapToStruct(rtc.claims, &out)

	return &out, nil
}
