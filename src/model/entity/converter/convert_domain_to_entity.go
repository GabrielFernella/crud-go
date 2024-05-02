package converter

import (
	"github.com/GabrielFernella/crud-go/src/model"
	"github.com/GabrielFernella/crud-go/src/model/entity"
)

func ConvertDomainToEntity(
	domain model.UserDomainInterface,
) *entity.UserEntity {
	return &entity.UserEntity{
		Email:    domain.GetEmail(),
		Password: domain.GetPassword(),
		Name:     domain.GetName(),
		Age:      domain.GetAge(),
	}
}
