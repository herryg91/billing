package usertoken_usecase

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type RefreshTokenClaim struct {
	claims      map[string]interface{}
	issuedTime  time.Time
	expiredTime time.Time
}

func (RefreshTokenClaim) New(claims map[string]interface{}, expiryInSecond int) RefreshTokenClaim {
	return RefreshTokenClaim{
		claims:      claims,
		issuedTime:  time.Now().UTC(),
		expiredTime: time.Now().Add(time.Second * time.Duration(expiryInSecond)).UTC()}
}

func (rtc RefreshTokenClaim) ToJwtClaims() jwt.MapClaims {
	jwtClaims := jwt.MapClaims{}
	for k, v := range rtc.claims {
		jwtClaims[k] = v
	}
	jwtClaims["issued_time"] = rtc.issuedTime.Unix()
	jwtClaims["expired_time"] = rtc.expiredTime.Unix()

	return jwtClaims
}

func (RefreshTokenClaim) FromJwtClaims(claims jwt.MapClaims) RefreshTokenClaim {
	rtc := RefreshTokenClaim{
		claims: map[string]interface{}{},
	}

	for k, v := range claims {
		if k == "issued_time" {
			rtc.issuedTime, _ = parseMapClaimTime(claims, k)
		} else if k == "expired_time" {
			rtc.expiredTime, _ = parseMapClaimTime(claims, k)
		} else {
			rtc.claims[k] = v
		}
	}

	// log.Println(rtc.expiredTime.Format(time.RFC3339))
	return rtc
}

func (rtc RefreshTokenClaim) IsExpired() bool {
	return rtc.expiredTime.UTC().Unix() < time.Now().UTC().Unix()
}
