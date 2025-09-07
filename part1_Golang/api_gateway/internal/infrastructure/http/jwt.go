package http

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("aP9x!7#sQ2vL8zY3mN1bR6kT0wE5hU4j")

type JWTClaims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func ValidateJWT(tokenString string) (string, string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecret, nil
	})
	if err != nil {
		return "", "", err
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok || !token.Valid {
		return "", "", errors.New("invalid token claims")
	}

	return claims.UserID, claims.Role, nil
}
