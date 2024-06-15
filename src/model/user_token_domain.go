package model

import (
	"fmt"
	"github.com/Vinicius-Madeira/go-web-app/src/configuration/rest_err"
	"github.com/golang-jwt/jwt"
	"os"
	"strings"
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

func VerifyToken(tokenValue string) (UserDomainInterface, *rest_err.RestError) {
	tokenString := RemoveBearerPrefix(tokenValue)
	secret := os.Getenv(jwtSecretKey)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(secret), nil
		}

		return nil, rest_err.NewBadRequestError("invalid token")
	})

	if err != nil {
		return nil, rest_err.NewUnauthorizedRequestError("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, rest_err.NewUnauthorizedRequestError("invalid token")
	}

	return &userDomain{
		id:       claims["id"].(string),
		email:    claims["email"].(string),
		password: claims["name"].(string),
		name:     claims["name"].(string),
		age:      int8(claims["age"].(float64)),
	}, nil
}

func RemoveBearerPrefix(token string) string {
	if strings.HasPrefix(token, "Bearer") {
		token = strings.TrimPrefix("Bearer ", token)
	}

	return token
}
