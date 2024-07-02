package model

import (
	"fmt"
	"github.com/Vinicius-Madeira/go-web-app/src/configuration/logger"
	"github.com/Vinicius-Madeira/go-web-app/src/configuration/rest_err"
	"github.com/gin-gonic/gin"
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

func VerifyTokenMiddleware(c *gin.Context) {

	secret := os.Getenv(jwtSecretKey)
	tokenValue := RemoveBearerPrefix(c.Request.Header.Get("Authorization"))

	token, err := jwt.Parse(tokenValue, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(secret), nil
		}

		return nil, rest_err.NewBadRequestError("invalid token")
	})

	if err != nil {
		restErr := rest_err.NewUnauthorizedRequestError("invalid token")
		c.JSON(restErr.Code, restErr)
		c.Abort()
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		restErr := rest_err.NewUnauthorizedRequestError("invalid token")
		c.JSON(restErr.Code, restErr)
		c.Abort()
		return
	}

	ud := &userDomain{
		id:       claims["id"].(string),
		email:    claims["email"].(string),
		password: claims["name"].(string),
		name:     claims["name"].(string),
		age:      int8(claims["age"].(float64)),
	}
	logger.Info(fmt.Sprintf("User authenticated: %#v", ud))

}

func RemoveBearerPrefix(token string) string {
	return strings.TrimPrefix("Bearer ", token)
}
