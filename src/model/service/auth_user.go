package service

import (
	"github.com/Vinicius-Madeira/go-web-app/src/configuration/logger"
	"github.com/Vinicius-Madeira/go-web-app/src/configuration/rest_err"
	"github.com/Vinicius-Madeira/go-web-app/src/model"
	"go.uber.org/zap"
)

func (uds *userDomainService) AuthUserServices(userDomain model.UserDomainInterface) (model.UserDomainInterface, string, *rest_err.RestError) {

	logger.Info("Init authUser model",
		zap.String("journey", "authUser"))

	userDomain.EncryptPassword()
	user, err := uds.findUserByEmailAndPasswordServices(userDomain.GetEmail(), userDomain.GetPassword())
	if err != nil {
		return nil, "", err
	}

	token, err := user.GenerateToken()
	if err != nil {
		return nil, "", err
	}

	logger.Info("AuthUser service executed successfully",
		zap.String("journey", "authUser"),
		zap.String("userID", user.GetID()))
	return user, token, nil
}
