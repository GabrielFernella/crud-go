package controller

import (
	"net/http"

	"github.com/GabrielFernella/crud-go/src/configuration/logger"
	"github.com/GabrielFernella/crud-go/src/configuration/validation"
	"github.com/GabrielFernella/crud-go/src/controller/model/request"
	"github.com/GabrielFernella/crud-go/src/model"
	"github.com/GabrielFernella/crud-go/src/model/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller",
		zap.String("journey", "createUser"),
	)
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "createUser"))
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	// vc pega o struct do model e depois passa por parametro no service
	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)

	// Instanciando o service
	service := service.NewUserDomainService()

	if err := service.CreateUser(domain); err != nil {
		c.JSON(err.Code, err)
		return
	}

	// response := response.UserResponse{
	// 	ID:    "teste",
	// 	Email: userRequest.Email,
	// 	Name:  userRequest.Name,
	// 	Age:   userRequest.Age,
	// }
	logger.Info("User created successfully",
		zap.String("journey", "createUser"))

	c.JSON(http.StatusOK, "")
}
