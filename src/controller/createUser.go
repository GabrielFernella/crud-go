package controller

import (
	"fmt"
	"net/http"

	"github.com/GabrielFernella/crud-go/src/configuration/validation"
	"github.com/GabrielFernella/crud-go/src/controller/model/request"
	"github.com/GabrielFernella/crud-go/src/controller/model/response"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		// restErr := rest_err.NewBadRequestError(fmt.Sprintf("There are some incorrect filds, error=%s\n", err.Error()))

		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	fmt.Println(userRequest)
	response := response.UserResponse{
		ID:    "test",
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	}

	c.JSON(http.StatusOK, response)

}
