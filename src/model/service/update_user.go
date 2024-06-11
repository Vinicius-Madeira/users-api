package service

import (
	"github.com/Vinicius-Madeira/go-web-app/src/configuration/rest_err"
	"github.com/Vinicius-Madeira/go-web-app/src/model"
)

func (*userDomainService) UpdateUser(userID string, userDomain model.UserDomainInterface) *rest_err.RestError {
	return nil
}
