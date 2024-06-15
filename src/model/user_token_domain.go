package model

import (
	"fmt"
	"github.com/Vinicius-Madeira/go-web-app/src/configuration/rest_err"
	"github.com/golang-jwt/jwt"
	"os"
	"time"
)

var (
	jwtSecretKey = "JWT_SECRET_KEY"
)

func (ud *userDomain) GenerateToken() (string, *rest_err.RestError) {
	secret := os.Getenv(jwtSecretKey)

	claims := jwt.MapClaims{
		"id":    ud.id,
		"email": ud.email,
		"name":  ud.name,
		"age":   ud.age,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", rest_err.NewInternalServerError(fmt.Sprintf("Error trying to generate JWT token, err=%s", err.Error()))
	}

	return tokenString, nil
}
