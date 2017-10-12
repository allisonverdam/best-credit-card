package models

import (
	"github.com/go-ozzo/ozzo-validation"
)

type Order struct {
	Price string `json:"price"`
}

func (m Order) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Price, validation.Required),
	)
}
