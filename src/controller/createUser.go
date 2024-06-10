package controller

import (
	"fmt"
	"github.com/Vinicius-Madeira/go-web-app/src/configuration/validation"
	"github.com/Vinicius-Madeira/go-web-app/src/controller/model/request"
	"github.com/gin-gonic/gin"
	"log"
)

func CreateUser(c *gin.Context) {

	log.Println("Init createUser controller")
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		log.Printf("Error trying to marshal object, error:%s\n", err.Error())
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userRequest)
}
