package model

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetAge() int8
	GetName() string

	GetID() string

	SetID(string)

	// GetJSONValue() (string, error)

	EncryptPassword()
}

func NewUserDomain(
	email, password, name string,
	age int8,
) UserDomainInterface {
	return &userDomain{
		email:    email,
		name:     name,
		age:      age,
		password: password,
	}
}
