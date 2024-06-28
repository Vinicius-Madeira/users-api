package repository

import (
	"context"
	"github.com/Vinicius-Madeira/go-web-app/src/configuration/logger"
	"github.com/Vinicius-Madeira/go-web-app/src/configuration/rest_err"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"os"
)

func (ur *userRepository) DeleteUser(userId string) *rest_err.RestError {
	logger.Info("Init deleteUser repository", zap.String("journey", "deleteUser"))

	collectionName := os.Getenv(mongodbCollection)
	collection := ur.databaseConnection.Collection(collectionName)

	userIdHex, _ := primitive.ObjectIDFromHex(userId)
	filter := bson.D{{Key: "_id", Value: userIdHex}}

	r, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		logger.Error("Error trying to delete user from database",
			err,
			zap.String("journey", "deleteUser"))
		return rest_err.NewInternalServerError(err.Error())
	}
	if r.DeletedCount == 0 {
		logger.Error("User not found in database",
			err,
			zap.String("journey", "deleteUser"))
		return rest_err.NewNotFoundError("User not found in database")
	}

	logger.Info("deleteUser repository executed successfully",
		zap.String("journey", "deleteUser"),
		zap.String("userID", userId))
	return nil
}
