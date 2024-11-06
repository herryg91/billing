package usertoken_usecase

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func parseMapClaimString(claims jwt.MapClaims, key string) (string, error) {
	result := ""
	if val, ok := claims[key]; ok {
		if valCasted, ok := val.(string); ok {
			result = valCasted
		} else {
			return "", fmt.Errorf("parse error: claims '%s' failed to parse: %v", key, val)
		}
	} else {
		return "", fmt.Errorf("invalid token: claims '%s' is missing", key)
	}
	return result, nil
}

func parseMapClaimInt(claims jwt.MapClaims, key string) (int, error) {
	result := -1
	if val, ok := claims[key]; ok {
		if valCasted, ok := val.(float64); ok {
			result = int(valCasted)
		} else {
			return -1, fmt.Errorf("parse error: claims '%s' failed to parse: %v", key, val)
		}
	} else {
		return -1, fmt.Errorf("invalid token: claims '%s' is missing", key)
	}
	return result, nil
}

func parseMapClaimTime(claims jwt.MapClaims, key string) (time.Time, error) {
	var result time.Time
	if val, ok := claims[key]; ok {
		if valCasted, ok := val.(float64); ok {
			result = time.Unix(int64(valCasted), 0)
		} else {
			return result, fmt.Errorf("parse error: claims '%s' failed to parse: %v", key, val)
		}
	} else {
		return result, fmt.Errorf("invalid token: claims '%s' is missing", key)
	}
	return result, nil
}

func generateToken(claims jwt.MapClaims, secretKey string) (token string, err error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = jwtToken.SignedString([]byte(secretKey))
	if err != nil {
		err = fmt.Errorf("<%w>: %s", ErrGenerateJwtToken, err.Error())
	}
	return
}
