package controller

import (
	"github.com/Vinicius-Madeira/go-web-app/src/configuration/logger"
	"github.com/Vinicius-Madeira/go-web-app/src/configuration/rest_err"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"net/http"
)

func (uc *userControllerInterface) DeleteUser(c *gin.Context) {

	logger.Info("Init deleteUser controller",
		zap.String("journey", "deleteUser"))

	userId := c.Param("userId")
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		restErr := rest_err.NewBadRequestError("Invalid userId, must be a hex value")
		c.JSON(restErr.Code, restErr)
		return
	}

	err := uc.service.DeleteUserServices(userId)
	if err != nil {
		logger.Error("Error trying to call deleteUser service", err, zap.String("journey", "deleteUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("deleteUser controller executed successfully",
		zap.String("journey", "deleteUser"),
		zap.String("userId", userId))

	c.Status(http.StatusNoContent)
}
