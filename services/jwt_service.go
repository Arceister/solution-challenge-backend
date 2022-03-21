package services

import (
	"api-solution/lib"
	"api-solution/models"
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

type JWTAuthService struct {
	env lib.Env
}

type UserToken struct {
	Email string
	Exp   uint
	ID    uint
	Name  string
	Phone string
}

func NewJWTAuthService(env lib.Env) JWTAuthService {
	return JWTAuthService{
		env: env,
	}
}

func (s JWTAuthService) ExtractClaims(tokenString string) jwt.MapClaims {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(s.env.JWTSecret), nil
	})

	if err != nil {
		return nil
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims
	} else {
		return nil
	}
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

func (s JWTAuthService) CreateToken(user models.User) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
		"phone": user.NoHP,
		"exp":   user.XPPoints,
		"lat":   user.Latitude,
		"lng":   user.Longitude,
	})

	tokenstring, err := token.SignedString([]byte(s.env.JWTSecret))

	if err != nil {
		fmt.Print("JWT Validation failed: " + err.Error())
	}
	return tokenstring
}
