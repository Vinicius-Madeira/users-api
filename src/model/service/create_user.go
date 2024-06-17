package service

import (
	"github.com/Vinicius-Madeira/go-web-app/src/configuration/logger"
	"github.com/Vinicius-Madeira/go-web-app/src/configuration/rest_err"
	"github.com/Vinicius-Madeira/go-web-app/src/model"
	"go.uber.org/zap"
)

func (uds *userDomainService) CreateUserServices(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestError) {

	logger.Info("Init createUser model",
		zap.String("journey", "createUser"))

	user, _ := uds.FindUserByEmailServices(userDomain.GetEmail())
	if user != nil {
		return nil, rest_err.NewBadRequestError("Email is already being used")
	}

	userDomain.EncryptPassword()

	userDomainRepository, err := uds.userRepository.CreateUser(userDomain)
	if err != nil {
		logger.Error("Error trying to call repository",
			err,
			zap.String("journey", "createUser"))
		return nil, err
	}

	logger.Info("CreateUser service executed successfully",
		zap.String("journey", "createUser"),
		zap.String("userID", userDomainRepository.GetID()))
	return userDomainRepository, nil
}
