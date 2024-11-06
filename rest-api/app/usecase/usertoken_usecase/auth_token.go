package usertoken_usecase

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type AuthTokenClaim struct {
	claims       map[string]interface{}
	RefreshToken string
	issuedTime   time.Time
	expiredTime  time.Time
}

func (AuthTokenClaim) New(claims map[string]interface{}, refreshToken string, expiryInSecond int) AuthTokenClaim {

	return AuthTokenClaim{
		claims:       claims,
		RefreshToken: refreshToken,
		issuedTime:   time.Now().UTC(),
		expiredTime:  time.Now().Add(time.Second * time.Duration(expiryInSecond)).UTC()}
}

func (atc AuthTokenClaim) ToJwtClaims() jwt.MapClaims {
	jwtClaims := jwt.MapClaims{}
	for k, v := range atc.claims {
		jwtClaims[k] = v
	}
	jwtClaims["refresh_token"] = atc.RefreshToken
	jwtClaims["issued_time"] = atc.issuedTime.Unix()
	jwtClaims["expired_time"] = atc.expiredTime.Unix()

	return jwtClaims
}

func (AuthTokenClaim) FromJwtClaims(claims jwt.MapClaims) AuthTokenClaim {
	atc := AuthTokenClaim{
		claims: map[string]interface{}{},
	}

	for k, v := range claims {
		if k == "refresh_token" {
			atc.RefreshToken, _ = parseMapClaimString(claims, k)
		} else if k == "issued_time" {
			atc.issuedTime, _ = parseMapClaimTime(claims, k)
		} else if k == "expired_time" {
			atc.expiredTime, _ = parseMapClaimTime(claims, k)
		} else {
			atc.claims[k] = v
		}
	}
	return atc
}

func (atc AuthTokenClaim) IsExpired() bool {
	return atc.expiredTime.UTC().Unix() < time.Now().UTC().Unix()
}
