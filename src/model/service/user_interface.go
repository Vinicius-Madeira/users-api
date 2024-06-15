package service

import (
	"github.com/Vinicius-Madeira/go-web-app/src/configuration/rest_err"
	"github.com/Vinicius-Madeira/go-web-app/src/model"
	"github.com/Vinicius-Madeira/go-web-app/src/model/repository"
)

func NewUserDomainService(userRepository repository.UserRepository) UserDomainService {
	return &userDomainService{
		userRepository: userRepository,
	}
}

type userDomainService struct {
	userRepository repository.UserRepository
}

type UserDomainService interface {
	CreateUserServices(domainInterface model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestError)
	FindUserByEmailServices(email string) (model.UserDomainInterface, *rest_err.RestError)
	FindUserByIDServices(id string) (model.UserDomainInterface, *rest_err.RestError)
	UpdateUserServices(string, model.UserDomainInterface) *rest_err.RestError
	DeleteUserServices(string) *rest_err.RestError
}
