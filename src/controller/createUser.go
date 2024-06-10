package controller

import (
	"github.com/Vinicius-Madeira/go-web-app/src/configuration/logger"
	"github.com/Vinicius-Madeira/go-web-app/src/configuration/validation"
	"github.com/Vinicius-Madeira/go-web-app/src/controller/model/request"
	"github.com/Vinicius-Madeira/go-web-app/src/controller/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func CreateUser(c *gin.Context) {

	logger.Info("Init CreateUser controller",
		zap.String("jorney", "createUser"))

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("jorney", "createUser"))
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	userResponse := response.UserResponse{
		ID:    "test",
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	}

	logger.Info("User created successfully",
		zap.String("jorney", "createUser"))

	c.JSON(http.StatusCreated, userResponse)
}
