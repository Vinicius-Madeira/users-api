package controller

import (
	"github.com/Vinicius-Madeira/go-web-app/src/configuration/logger"
	"github.com/Vinicius-Madeira/go-web-app/src/configuration/rest_err"
	"github.com/Vinicius-Madeira/go-web-app/src/view"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"net/http"
	"net/mail"
)

// FindUserById retrieves user information based on the provided user ID.
// @Summary Find User by ID
// @Description Retrieves user details based on the user ID provided as a parameter.
// @Tags Users
// @Accept json
// @Produce json
// @Param userId path string true "ID of the user to be retrieved"
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {object} response.UserResponse "User information retrieved successfully"
// @Failure 400 {object} rest_err.RestError "Error: Invalid user ID"
// @Failure 404 {object} rest_err.RestError "User not found"
// @Router /getUserById/{userId} [get]
func (uc *userControllerInterface) FindUserById(c *gin.Context) {
	logger.Info("Init FindUserById controller",
		zap.String("journey", "findUserByID"))

	userID := c.Param("userId")
	if _, err := primitive.ObjectIDFromHex(userID); err != nil {

		logger.Error("Error trying to validate userID", err,
			zap.String("journey", "findUserByID"),
			zap.String("userID", userID))

		errorMessage := "UserID is not a valid id"

		c.JSON(http.StatusBadRequest, rest_err.NewBadRequestError(errorMessage))
		return
	}

	userDomain, err := uc.service.FindUserByIDServices(userID)
	if err != nil {
		logger.Error("Error trying to call findUserByID services", err,
			zap.String("journey", "findUserByID"),
			zap.String("userID", userID))
		c.JSON(err.Code, err)
		return
	}

	userResponse := view.ConvertDomainToResponse(userDomain)
	logger.Info("FindUserByID controller executed successfully",
		zap.String("journey", "findUserByID"),
		zap.String("userID", userID))

	c.JSON(http.StatusOK, userResponse)
}

// FindUserByEmail retrieves user information based on the provided email.
// @Summary Find User by Email
// @Description Retrieves user details based on the email provided as a parameter.
// @Tags Users
// @Accept json
// @Produce json
// @Param userEmail path string true "Email of the user to be retrieved"
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {object} response.UserResponse "User information retrieved successfully"
// @Failure 400 {object} rest_err.RestError "Error: Invalid user ID"
// @Failure 404 {object} rest_err.RestError "User not found"
// @Router /getUserByEmail/{userEmail} [get]
func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {
	logger.Info("Init FindUserByEmail controller", zap.String("journey", "findUserByEmail"))

	userEmail := c.Param("userEmail")
	if _, err := mail.ParseAddress(userEmail); err != nil {

		logger.Error("Error trying to validate userEmail", err,
			zap.String("journey", "findUserByEmail"),
			zap.String("userEmail", userEmail))

		errorMessage := "UserEmail is not a valid userEmail"

		c.JSON(http.StatusBadRequest, rest_err.NewBadRequestError(errorMessage))
		return
	}

	userDomain, err := uc.service.FindUserByEmailServices(userEmail)
	if err != nil {
		logger.Error("Error trying to call FindUserByEmail services", err,
			zap.String("journey", "findUserByEmail"),
			zap.String("userEmail", userEmail))
		c.JSON(err.Code, err)
		return
	}

	userResponse := view.ConvertDomainToResponse(userDomain)
	logger.Info("FindUserByEmail controller executed successfully",
		zap.String("journey", "findUserByEmail"),
		zap.String("userEmail", userEmail))

	c.JSON(http.StatusOK, userResponse)
}
