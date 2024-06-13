package view

import (
	"github.com/Vinicius-Madeira/go-web-app/src/controller/model/response"
	"github.com/Vinicius-Madeira/go-web-app/src/model"
)

func ConvertDomainToResponse(
	userDomain model.UserDomainInterface,
) response.UserResponse {
	return response.UserResponse{
		ID:    userDomain.GetID(),
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   userDomain.GetAge(),
	}
}
