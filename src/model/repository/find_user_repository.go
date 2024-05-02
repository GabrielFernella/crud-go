package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/GabrielFernella/crud-go/src/configuration/logger"
	"github.com/GabrielFernella/crud-go/src/configuration/rest_err"
	"github.com/GabrielFernella/crud-go/src/model"
	"github.com/GabrielFernella/crud-go/src/model/entity"
	"github.com/GabrielFernella/crud-go/src/model/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func (ur *userRepository) FindUserByEmail(email string) (model.UserDomainInterface, *rest_err.RestErr) {

	logger.Info("Init findUserByEmail user repository", zap.String("journey", "findUserByEmail"))

	collection_name := os.Getenv(MONGODB_USER_COLLECTION)

	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}

	filter := bson.D{{Key: "email", Value: email}}

	err := collection.FindOne(context.Background(), filter).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("User not found with this email: %s", email)
			logger.Error(errorMessage, err, zap.String("journey", "findUserByEmail"))

			return nil, rest_err.NewNotFoundError(errorMessage)
		}

		errorMessage := "Error trying to find user by email"
		logger.Error(errorMessage, err, zap.String("journey", "findUserByEmail"))
		return nil, rest_err.NewInternalServerError(errorMessage)
	}

	logger.Info("FindUserByEmail repository executed successfully", zap.String("journey", "findUserByEmail"), zap.String("email", email), zap.String("id", userEntity.ID.Hex()))

	return converter.ConvertEntityToDomain(*userEntity), nil
}

func (ur *userRepository) FindUserByID(id string) (model.UserDomainInterface, *rest_err.RestErr) {

	logger.Info("Init findUserById user repository", zap.String("journey", "findUserById"))

	collection_name := os.Getenv(MONGODB_USER_COLLECTION)

	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}

	filter := bson.D{{Key: "_id", Value: id}}

	err := collection.FindOne(context.Background(), filter).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("User not found with this id: %s", id)
			logger.Error(errorMessage, err, zap.String("journey", "findUserById"))

			return nil, rest_err.NewNotFoundError(errorMessage)
		}

		errorMessage := "Error trying to find user by id"
		logger.Error(errorMessage, err, zap.String("journey", "findUserById"))
		return nil, rest_err.NewInternalServerError(errorMessage)
	}

	logger.Info("FindUserById repository executed successfully", zap.String("journey", "findUserById"), zap.String("userId", userEntity.ID.Hex()))

	return converter.ConvertEntityToDomain(*userEntity), nil
}
