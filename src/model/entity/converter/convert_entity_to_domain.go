package converter

import (
	"github.com/GabrielFernella/crud-go/src/model"
	"github.com/GabrielFernella/crud-go/src/model/entity"
)

func ConvertEntityToDomain(
	entity entity.UserEntity,
) model.UserDomainInterface {
	domain := model.NewUserDomain(
		entity.Email,
		entity.Password,
		entity.Name,
		entity.Age,
	)

	domain.SetID(entity.ID.Hex())

	return domain
}
