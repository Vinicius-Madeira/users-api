package controller

import (
	"fmt"
	"github.com/Vinicius-Madeira/go-web-app/src/configuration/logger"
	"github.com/Vinicius-Madeira/go-web-app/src/configuration/rest_err"
	"github.com/Vinicius-Madeira/go-web-app/src/model"
	"github.com/Vinicius-Madeira/go-web-app/src/view"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"net/http"
	"net/mail"
)

func (uc *userControllerInterface) FindUserById(c *gin.Context) {
	logger.Info("Init FindUserById controller",
		zap.String("journey", "findUserByID"))

	user, err := model.VerifyToken(c.Request.Header.Get("Authorization"))
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	logger.Info(fmt.Sprintf("User authenticated: %#v", user))

	userID := c.Param("userId")
	if _, err := primitive.ObjectIDFromHex(userID); err != nil {

		logger.Error("Error trying to validate userID", err,
			zap.String("journey", "findUserByID"),
			zap.String("userID", userID))

		errorMessage := "UserID is not a valid id"

		c.JSON(http.StatusBadRequest, rest_err.NewBadRequestError(errorMessage))
		return
	}

	userDomain, err := uc.service.FindUserByIDServices(userID)
	if err != nil {
		logger.Error("Error trying to call findUserByID services", err,
			zap.String("journey", "findUserByID"),
			zap.String("userID", userID))
		c.JSON(err.Code, err)
		return
	}

	userResponse := view.ConvertDomainToResponse(userDomain)
	logger.Info("FindUserByID controller executed successfully",
		zap.String("journey", "findUserByID"),
		zap.String("userID", userID))

	c.JSON(http.StatusOK, userResponse)
}

func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {
	logger.Info("Init FindUserByEmail controller", zap.String("journey", "findUserByEmail"))

	user, err := model.VerifyToken(c.Request.Header.Get("Authorization"))
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	logger.Info(fmt.Sprintf("User authenticated: %#v", user))

	userEmail := c.Param("userEmail")
	if _, err := mail.ParseAddress(userEmail); err != nil {

		logger.Error("Error trying to validate userEmail", err,
			zap.String("journey", "findUserByEmail"),
			zap.String("userEmail", userEmail))

		errorMessage := "UserEmail is not a valid userEmail"

		c.JSON(http.StatusBadRequest, rest_err.NewBadRequestError(errorMessage))
		return
	}

	userDomain, err := uc.service.FindUserByEmailServices(userEmail)
	if err != nil {
		logger.Error("Error trying to call FindUserByEmail services", err,
			zap.String("journey", "findUserByEmail"),
			zap.String("userEmail", userEmail))
		c.JSON(err.Code, err)
		return
	}

	userResponse := view.ConvertDomainToResponse(userDomain)
	logger.Info("FindUserByEmail controller executed successfully",
		zap.String("journey", "findUserByEmail"),
		zap.String("userEmail", userEmail))

	c.JSON(http.StatusOK, userResponse)
}
