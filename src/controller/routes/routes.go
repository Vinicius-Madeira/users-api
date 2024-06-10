package routes

import (
	"github.com/Vinicius-Madeira/go-web-app/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/getUserById/:userId", controller.FindUserById)          // get user by id
	r.GET("/getUserByEmail/:userEmail", controller.FindUserByEmail) // get user by email
	r.POST("/createUser/", controller.CreateUser)
	r.PUT("/updateUser/:userId", controller.UpdateUser)
	r.DELETE("/deleteUser/:userId", controller.DeleteUser)
}
