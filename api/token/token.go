package token

import (
	"api-gateway/config"
	"fmt"

	"github.com/golang-jwt/jwt"
)

func ValidateToken(tokenstr string) (bool, error) {
	_, err := ExtractClaims(tokenstr)
	if err != nil {
		return false, err
	}
	return true, nil
}

func ExtractClaims(tokenstr string) (jwt.MapClaims, error) {
	token, err := jwt.ParseWithClaims(tokenstr, jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(config.Load().SIGNING_KEY), nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}
	return claims, nil
}
