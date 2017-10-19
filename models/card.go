package models

import (
	"github.com/go-ozzo/ozzo-validation"
)

type Card struct {
	Id              int     `json:"id" db:"id"`
	Number          string  `json:"number" db:"cc_number"`
	DueDate         int     `json:"due_date" db:"cc_due_date"`
	ExpirationMonth int     `json:"expiration_month" db:"cc_expiration_month"`
	ExpirationYear  int     `json:"expiration_year" db:"cc_expiration_year"`
	CVV             int     `json:"cvv" db:"cc_cvv"`
	Limit           float64 `json:"limit" db:"cc_limit"`
	WalletId        int     `json:"wallet_id" db:"wallet_id"`
}

func (m Card) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Number, validation.Required, validation.Length(16, 16)),
	)
}
