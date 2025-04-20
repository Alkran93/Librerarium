package util

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
)

var JwtSecret = []byte("super-secret-key")

func GetUsernameFromToken(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return "", errors.New("missing Bearer token")
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return JwtSecret, nil
	})
	if err != nil || !token.Valid {
		return "", errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid claims")
	}

	username, ok := claims["sub"].(string)
	if !ok {
		return "", errors.New("no sub in token")
	}

	return username, nil
}
