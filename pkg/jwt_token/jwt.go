package jwt_token

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func GenerateJwtToken(id int64, username, secreteKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       id,
		"username": username,
		"exp":      time.Now().Add(30 * time.Hour),
	})

	key := []byte(secreteKey)
	signedToken, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func ValidateJwtToken(tokenStr, secreteKey string) (int64, string, error) {
	claims := jwt.MapClaims{}
	key := []byte(secreteKey)
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		return 0, "", err
	}

	if !token.Valid {
		return 0, "", errors.New("invalid token")
	}

	return int64(claims["id"].(float64)), claims["username"].(string), nil
}
