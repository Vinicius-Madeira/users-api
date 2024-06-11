package model

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/Vinicius-Madeira/go-web-app/src/configuration/rest_err"
)

func NewUserDomain(
	email, password, name string,
	age int8,
) UserDomainInterface {
	return &UserDomain{
		email, password, name, age,
	}
}

type UserDomain struct {
	Email    string
	Password string
	Name     string
	Age      int8
}

type UserDomainInterface interface {
	CreateUser() *rest_err.RestError
	UpdateUser(string) *rest_err.RestError
	FindUser(string) (*UserDomain, *rest_err.RestError)
	DeleteUser(string) *rest_err.RestError
}

func (ud *UserDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.Password))
	ud.Password = hex.EncodeToString(hash.Sum(nil))
}
