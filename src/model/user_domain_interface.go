package model

import "github.com/Vinicius-Madeira/go-web-app/src/configuration/rest_err"

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetAge() int8
	GetName() string
	GetID() string
	SetID(string)
	EncryptPassword()

	GenerateToken() (string, *rest_err.RestError)
}

func NewUserDomain(
	email, password, name string,
	age int8,
) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: password,
		name:     name,
		age:      age,
	}
}

func NewUserAuthDomain(
	email, password string,
) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: password,
	}
}

func NewUserUpdateDomain(
	name string,
	age int8,
) UserDomainInterface {
	return &userDomain{
		name: name,
		age:  age,
	}
}
