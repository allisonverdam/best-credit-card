package models

import (
	"github.com/go-ozzo/ozzo-validation"
)

type Order struct {
	Price    float64 `json:"price" description:"Valor da transação"`
	WalletId int     `json:"wallet_id" description:"ID da carteira que será utilizado nessa transação."`
	CardId   int     `json:"card_id" description:"ID do cartão que será utilizado nessa transação."`
}

func (m Order) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Price, validation.Required, validation.Min(0.0)),
		validation.Field(&m.WalletId, validation.Required),
	)
}

func (m Order) ValidateCardIdAndPrice() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.CardId, validation.Required),
		validation.Field(&m.Price, validation.Required, validation.Min(0.0)),
	)
}
