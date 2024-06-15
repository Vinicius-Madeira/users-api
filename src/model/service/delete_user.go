package service

import (
	"github.com/Vinicius-Madeira/go-web-app/src/configuration/logger"
	"github.com/Vinicius-Madeira/go-web-app/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (uds *userDomainService) DeleteUserServices(userId string) *rest_err.RestError {
	logger.Info("Init deleteUser model", zap.String("journey", "deleteUser"))

	err := uds.userRepository.DeleteUser(userId)
	if err != nil {
		logger.Error("Error trying to call repository", err, zap.String("journey", "deleteUser"))
		return err
	}

	logger.Info("DeleteUserServices service executed successfully",
		zap.String("journey", "deleteUser"),
		zap.String("userID", userId))
	return nil
}
