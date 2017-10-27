package models

import (
	"github.com/go-ozzo/ozzo-validation"
)

type Wallet struct {
	Id           int     `json:"id" db:"id"`
	RealLimit    float64 `json:"real_limit" db:"real_limit"`
	MaximumLimit float64 `json:"maximum_limit" db:"maximum_limit"`
	PersonId     int     `json:"person_id" db:"person_id"`
}

func (p Wallet) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.RealLimit, validation.Required),
		validation.Field(&p.MaximumLimit, validation.Required),
	)
}
