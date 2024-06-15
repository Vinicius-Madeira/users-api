package routes

import (
	"github.com/Vinicius-Madeira/go-web-app/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup, userController controller.UserControllerInterface) {

	r.GET("/getUserById/:userId", userController.FindUserById)          // get user by id
	r.GET("/getUserByEmail/:userEmail", userController.FindUserByEmail) // get user by email
	r.POST("/createUser/", userController.CreateUser)
	r.PUT("/updateUser/:userId", userController.UpdateUser)
	r.DELETE("/deleteUser/:userId", userController.DeleteUser)
	r.POST("/auth", userController.AuthUser)
}
