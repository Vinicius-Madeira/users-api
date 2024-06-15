package service

import (
	"github.com/Vinicius-Madeira/go-web-app/src/configuration/logger"
	"github.com/Vinicius-Madeira/go-web-app/src/configuration/rest_err"
	"github.com/Vinicius-Madeira/go-web-app/src/model"
	"go.uber.org/zap"
)

func (uds *userDomainService) FindUserByIDServices(
	id string,
) (model.UserDomainInterface, *rest_err.RestError) {

	logger.Info("Init findUserByID services.",
		zap.String("journey", "findUserByID"))

	return uds.userRepository.FindUserByID(id)
}

func (uds *userDomainService) FindUserByEmailServices(
	email string,
) (model.UserDomainInterface, *rest_err.RestError) {

	logger.Info("Init findUserByEmail services.",
		zap.String("journey", "findUserByEmail"))

	return uds.userRepository.FindUserByEmail(email)
}

func (uds *userDomainService) findUserByEmailAndPasswordServices(
	email string,
	password string,
) (model.UserDomainInterface, *rest_err.RestError) {

	logger.Info("Init findUserByEmail services.",
		zap.String("journey", "findUserByEmail"))

	return uds.userRepository.FindUserByEmailAndPassword(email, password)
}
