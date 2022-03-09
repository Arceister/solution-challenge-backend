package services

import (
	"api-solution/lib"
	"errors"

	"github.com/golang-jwt/jwt/v4"
)

type JWTAuthService struct {
	env lib.Env
}

func (s JWTAuthService) Authorize(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(s.env.JWTSecret), nil
	})
	if token.Valid {
		return true, nil
	} else if verification, ok := err.(*jwt.ValidationError); ok {
		if verification.Errors&jwt.ValidationErrorMalformed != 0 {
			return false, errors.New("token malformed")
		}
		if verification.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return false, errors.New("token expired")
		}
	}
	return false, errors.New("cannot handle token")
}
