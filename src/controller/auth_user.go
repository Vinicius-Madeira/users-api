package controller

import (
	"github.com/Vinicius-Madeira/go-web-app/src/configuration/logger"
	"github.com/Vinicius-Madeira/go-web-app/src/configuration/validation"
	"github.com/Vinicius-Madeira/go-web-app/src/controller/model/request"
	"github.com/Vinicius-Madeira/go-web-app/src/model"
	"github.com/Vinicius-Madeira/go-web-app/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func (uc *userControllerInterface) AuthUser(c *gin.Context) {
	logger.Info("Init AuthUser controller",
		zap.String("journey", "authUser"))

	var userRequest request.UserAuth

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "authUser"))
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}
	domain := model.NewUserAuthDomain(userRequest.Email, userRequest.Password)

	domainResult, err := uc.service.AuthUserServices(domain)
	if err != nil {
		logger.Error("Error trying to call AuthUser service", err, zap.String("journey", "authUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("AuthUser controller executed successfully",
		zap.String("journey", "authUser"),
		zap.String("userID", domainResult.GetID()))

	c.JSON(http.StatusCreated, view.ConvertDomainToResponse(domainResult))
}
