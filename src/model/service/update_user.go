package service

import (
	"github.com/GabrielFernella/crud-go/src/configuration/rest_err"
	"github.com/GabrielFernella/crud-go/src/model"
)

func (*userDomainService) UpdateUser(
	userId string,
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {
	return nil
}
