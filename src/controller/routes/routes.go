package routes

import (
	"github.com/Vinicius-Madeira/go-web-app/src/controller"
	"github.com/Vinicius-Madeira/go-web-app/src/model"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup, userController controller.UserControllerInterface) {

	r.GET("/getUserById/:userId", model.VerifyTokenMiddleware, userController.FindUserById)          // get user by id
	r.GET("/getUserByEmail/:userEmail", model.VerifyTokenMiddleware, userController.FindUserByEmail) // get user by email
	r.POST("/createUser/", userController.CreateUser)                                                // create a new user
	r.PUT("/updateUser/:userId", model.VerifyTokenMiddleware, userController.UpdateUser)             // updates a user
	r.DELETE("/deleteUser/:userId", model.VerifyTokenMiddleware, userController.DeleteUser)          // deletes a user
	r.POST("/auth", userController.AuthUser)                                                         // authenticates the user
}
