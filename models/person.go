package models

import (
	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Person struct {
	Id       int    `json:"id" db:"id" description:"Identificador da pessoa."`
	Name     string `json:"name" db:"name" description:"Nome completo."`
	Username string `json:"username" db:"username" description:"Usuario."`
	Email    string `json:"email" db:"email" description:"Email."`
	Password string `json:"password,omitempty" db:"password" description:"Senha."`
}

func (p Person) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Name, validation.Required),
		validation.Field(&p.Username, validation.Required, validation.Length(0, 20)),
		validation.Field(&p.Email, validation.Required, is.Email),
		validation.Field(&p.Password, validation.Required, validation.Length(8, 255)),
	)
}

func (p Person) ValidateOnUpdate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Name, validation.Required),
		validation.Field(&p.Username, validation.Required, validation.Length(0, 20)),
		validation.Field(&p.Email, validation.Required, is.Email),
	)
}
