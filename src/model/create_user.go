package model

import (
	"fmt"
	"github.com/Vinicius-Madeira/go-web-app/src/configuration/logger"
	"github.com/Vinicius-Madeira/go-web-app/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *UserDomain) CreateUser() *rest_err.RestError {

	logger.Info("Init createUser model", zap.String("journey", "createUser"))

	ud.EncryptPassword()

	fmt.Println(ud)

	return nil
}
