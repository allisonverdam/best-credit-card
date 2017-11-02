package models

import (
	"github.com/go-ozzo/ozzo-validation"
)

type Card struct {
	Id              int     `json:"id" db:"id" description:"Identificador do cartão."`
	Number          string  `json:"number" db:"cc_number" description:"Número do cartão."`
	DueDate         int     `json:"due_date" db:"cc_due_date" description:"Data de vencimento."`
	ExpirationMonth int     `json:"expiration_month" db:"cc_expiration_month" description:"Mês de expiração."`
	ExpirationYear  int     `json:"expiration_year" db:"cc_expiration_year" description:"Ano de expiração."`
	CVV             int     `json:"cvv" db:"cc_cvv" description:"Código de verificação do cartão."`
	RealLimit       float64 `json:"real_limit" db:"cc_real_limit" description:"Limite real do cartão."`
	AvaliableLimit  float64 `json:"avaliable_limit" db:"cc_avaliable_limit" description:"Limite disponível."`
	Currency        string  `json:"currency" db:"cc_currency" description:"Tipo de moeda."`
	WalletId        int     `json:"wallet_id" db:"wallet_id" description:"ID da carteira que esse cartão pertence."`
}

func (m Card) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Number, validation.Required, validation.Length(16, 16)),
		validation.Field(&m.DueDate, validation.Required),
		validation.Field(&m.ExpirationMonth, validation.Required),
		validation.Field(&m.ExpirationYear, validation.Required),
		validation.Field(&m.CVV, validation.Required),
		validation.Field(&m.RealLimit, validation.Required),
		validation.Field(&m.Currency, validation.Required),
		validation.Field(&m.AvaliableLimit, validation.Required, validation.Max(float64(*&m.RealLimit))),
		validation.Field(&m.WalletId, validation.Required),
	)
}
