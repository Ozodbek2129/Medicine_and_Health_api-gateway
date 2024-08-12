package auth

import (
	"github.com/dgrijalva/jwt-go"
)

const (
	SIGNING_KEY = "R)#@&&*@^B"
)

func ValidateAccessToken(tokenStr string) (bool, error) {
	_, err := ExtractAccessClaim(tokenStr)
	if err != nil {
		return false, err
	}
	return true, nil
}

func ExtractAccessClaim(tokenStr string) (*jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(SIGNING_KEY), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		return nil, err
	}

	return &claims, nil
}

func GetUserInfoFromAccessToken(accessTokenString string) (string, string, string, string, string, string, string, error) {
	refreshToken, err := jwt.Parse(accessTokenString, func(token *jwt.Token) (interface{}, error) { return []byte(SIGNING_KEY), nil })
	if err != nil || !refreshToken.Valid {
		return "", "", "", "", "", "", "", err
	}
	claims, ok := refreshToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", "", "", "", "", "", "", err
	}
	userID := claims["user_id"].(string)
	Role := claims["role"].(string)
	Firstname := claims["first_name"].(string)
	Email := claims["email"].(string)
	DateOfBirth := claims["date_of_birth"].(string)
	LastName := claims["last_name"].(string)
	Gender := claims["gender"].(string)

	return userID, Role, Firstname, Email, DateOfBirth, LastName, Gender, nil
}
