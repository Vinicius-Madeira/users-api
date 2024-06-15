package service

import (
	"github.com/Vinicius-Madeira/go-web-app/src/configuration/logger"
	"github.com/Vinicius-Madeira/go-web-app/src/configuration/rest_err"
	"github.com/Vinicius-Madeira/go-web-app/src/model"
	"go.uber.org/zap"
)

func (uds *userDomainService) UpdateUserServices(userID string, userDomain model.UserDomainInterface) *rest_err.RestError {
	logger.Info("Init updateUser model", zap.String("journey", "updateUser"))

	err := uds.userRepository.UpdateUser(userID, userDomain)
	if err != nil {
		logger.Error("Error trying to call repository", err, zap.String("journey", "updateUser"))
		return err
	}

	logger.Info("UpdateUserServices service executed successfully",
		zap.String("journey", "updateUser"),
		zap.String("userID", userID))
	return nil
}
