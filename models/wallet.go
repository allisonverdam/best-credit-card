package models

import (
	"github.com/go-ozzo/ozzo-validation"
)

type Wallet struct {
	Id             int     `json:"id" db:"id" description:"Identificador da carteira."`
	CurrentLimit   float64 `json:"current_limit" db:"current_limit" description:"Limite atual da carteira."`
	MaximumLimit   float64 `json:"maximum_limit" db:"maximum_limit" description:"Limite maximo da carteira, é a soma dos limites dos cartoes."`
	AvaliableLimit float64 `json:"avaliable_limit" db:"avaliable_limit" description:"Limite disponível para upgrade. É igual a limite maximo - limite atual."`
	PersonId       int     `json:"person_id" db:"person_id" description:"ID da pessoa dona dessa carteira"`
}

func (p Wallet) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.PersonId, validation.Required),
	)
}
