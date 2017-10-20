package models

import (
	"github.com/go-ozzo/ozzo-validation"
)

type Order struct {
	Price    float64 `json:"price"`
	WalletId int     `json:"wallet_id"`
	CardId   int     `json:"card_id"`
}

func (m Order) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Price, validation.Required),
		validation.Field(&m.WalletId, validation.Required),
	)
}

func (m Order) ValidateCardIdAndPrice() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.CardId, validation.Required),
		validation.Field(&m.Price, validation.Required, validation.Min(0.0)),
	)
}
