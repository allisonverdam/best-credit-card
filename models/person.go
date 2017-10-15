package models

import (
	"github.com/go-ozzo/ozzo-validation"
)

type Person struct {
	Id       int    `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

func (p Person) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Username, validation.Required, validation.Length(0, 20)),
		validation.Field(&p.Password, validation.Required, validation.Length(8, 255)),
	)
}
