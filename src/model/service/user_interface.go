package service

import (
	"github.com/Vinicius-Madeira/go-web-app/src/configuration/rest_err"
	"github.com/Vinicius-Madeira/go-web-app/src/model"
)

func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}

type userDomainService struct {
}

type UserDomainService interface {
	CreateUser(domainInterface model.UserDomainInterface) *rest_err.RestError
	UpdateUser(string, model.UserDomainInterface) *rest_err.RestError
	FindUser(string) (*model.UserDomainInterface, *rest_err.RestError)
	DeleteUser(string) *rest_err.RestError
}
