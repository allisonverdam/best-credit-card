package services

import (
	"testing"

	"github.com/allisonverdam/best-credit-card/app"
	"github.com/allisonverdam/best-credit-card/daos"
	"github.com/allisonverdam/best-credit-card/models"
	"github.com/allisonverdam/best-credit-card/testdata"
	routing "github.com/go-ozzo/ozzo-routing"
	"github.com/stretchr/testify/assert"
)

func TestGetCard(t *testing.T) {
	db := testdata.ResetDB()
	dao := daos.NewCardDAO()
	service := NewCardService(dao)

	testDBCall(db, func(rs app.RequestScope, c routing.Context) {
		card, err := service.GetCard(rs, 1)
		assert.Nil(t, err)
		if assert.NotNil(t, card) {
			assert.Equal(t, "1234123412341230", card.Number)
		}
	})

}

func TestGetCardWithErrorCardNotFound(t *testing.T) {
	db := testdata.ResetDB()
	dao := daos.NewCardDAO()
	service := NewCardService(dao)

	testDBCall(db, func(rs app.RequestScope, c routing.Context) {
		card, err := service.GetCard(rs, 0)
		assert.Nil(t, card)
		if assert.NotNil(t, err) {
			assert.Equal(t, "sql: no rows in result set", err.Error())
		}

	})

}

func TestGetCardWithErrorCardDoesNotBelongToTheAuthenticatedUser(t *testing.T) {
	db := testdata.ResetDB()
	dao := daos.NewCardDAO()
	service := NewCardService(dao)

	testDBCall(db, func(rs app.RequestScope, c routing.Context) {
		card, err := service.GetCard(rs, 4)
		assert.Nil(t, card)
		if assert.NotNil(t, err) {
			assert.Equal(t, "You're not allowed to do this.", err.Error())
		}
	})

}

func TestPayCreditCard(t *testing.T) {
	db := testdata.ResetDB()
	dao := daos.NewCardDAO()
	service := NewCardService(dao)
	order := models.Order{
		Price:  100,
		CardId: 1,
	}

	testDBCall(db, func(rs app.RequestScope, c routing.Context) {
		card, err := service.PayCreditCard(rs, order)
		assert.Nil(t, err)
		//o card tinha 180 de limite disponivel
		if assert.NotNil(t, card) {
			assert.Equal(t, 280.0, card.AvaliableLimit)
		}
	})

}

func TestPayCreditCardWithErrorCardNotFound(t *testing.T) {
	db := testdata.ResetDB()
	dao := daos.NewCardDAO()
	service := NewCardService(dao)
	order := models.Order{
		Price:  100,
		CardId: 0,
	}

	testDBCall(db, func(rs app.RequestScope, c routing.Context) {
		card, err := service.PayCreditCard(rs, order)
		assert.Nil(t, card)
		if assert.NotNil(t, err) {
			assert.Equal(t, "sql: no rows in result set", err.Error())
		}
	})

}

func TestPayCreditCardWithErrorCardNotBelongToThisPerson(t *testing.T) {
	db := testdata.ResetDB()
	dao := daos.NewCardDAO()
	service := NewCardService(dao)
	order := models.Order{
		Price:  100,
		CardId: 4,
	}

	testDBCall(db, func(rs app.RequestScope, c routing.Context) {
		card, err := service.PayCreditCard(rs, order)
		assert.Nil(t, card)
		if assert.NotNil(t, err) {
			assert.Equal(t, "You're not allowed to do this.", err.Error())
		}
	})

}

func TestPayCreditCardWithErrorPriceGreaterThanCardLimit(t *testing.T) {
	db := testdata.ResetDB()
	dao := daos.NewCardDAO()
	service := NewCardService(dao)
	order := models.Order{
		Price:  10000,
		CardId: 1,
	}

	testDBCall(db, func(rs app.RequestScope, c routing.Context) {
		card, err := service.PayCreditCard(rs, order)
		assert.Nil(t, card)
		if assert.NotNil(t, err) {
			assert.Equal(t, "Price greater than the maximum card limit.", err.Error())
		}
	})

}

func TestGetBestCards(t *testing.T) {
	db := testdata.ResetDB()
	dao := daos.NewCardDAO()
	service := NewCardService(dao)
	order := models.Order{
		Price:    613,
		WalletId: 1,
	}

	testDBCall(db, func(rs app.RequestScope, c routing.Context) {
		cards, err := service.GetBestCards(rs, &order)
		assert.Nil(t, err)
		if assert.NotNil(t, cards) {
			assert.Equal(t, 2, len(*cards))
		}
	})

}

func TestGetBestCardsWithErrorLimitNotAvaliable(t *testing.T) {
	db := testdata.ResetDB()
	dao := daos.NewCardDAO()
	service := NewCardService(dao)
	order := models.Order{
		Price:    5000,
		WalletId: 1,
	}

	testDBCall(db, func(rs app.RequestScope, c routing.Context) {
		cards, err := service.GetBestCards(rs, &order)
		assert.Nil(t, cards)
		if assert.NotNil(t, err) {
			assert.Equal(t, "You don't have enough credit to make this purchase.", err.Error())
		}
	})

}

func TestGetCardsByWalletId(t *testing.T) {
	db := testdata.ResetDB()
	dao := daos.NewCardDAO()
	service := NewCardService(dao)

	testDBCall(db, func(rs app.RequestScope, c routing.Context) {
		cards, err := service.GetCardsByWalletId(rs, 1)
		assert.Nil(t, err)
		if assert.NotNil(t, cards) {
			assert.Equal(t, 3, len(*cards))
		}
	})

}

func TestGetCardsByWalletIdWithErrorWalletDontBelongToAuthenticatedPerson(t *testing.T) {
	db := testdata.ResetDB()
	dao := daos.NewCardDAO()
	service := NewCardService(dao)

	testDBCall(db, func(rs app.RequestScope, c routing.Context) {
		cards, err := service.GetCardsByWalletId(rs, 2)
		assert.Nil(t, cards)
		if assert.NotNil(t, err) {
			assert.Equal(t, "You're not allowed to do this.", err.Error())
		}
	})

}

func TestCreateCard(t *testing.T) {
	db := testdata.ResetDB()
	dao := daos.NewCardDAO()
	service := NewCardService(dao)
	tempCard := models.Card{
		AvaliableLimit:  100,
		Currency:        "BRL",
		CVV:             123,
		DueDate:         2,
		ExpirationMonth: 2,
		ExpirationYear:  22,
		Number:          "2122123233233122",
		RealLimit:       322,
		WalletId:        1,
	}

	testDBCall(db, func(rs app.RequestScope, c routing.Context) {
		card, err := service.CreateCard(rs, &tempCard)
		assert.Nil(t, err)
		if assert.NotNil(t, card) {
			assert.Equal(t, "2122123233233122", card.Number)
		}
	})

}

func TestCreateCardWithErrorWalletNotBelongToAuthenticatedPerson(t *testing.T) {
	db := testdata.ResetDB()
	dao := daos.NewCardDAO()
	service := NewCardService(dao)
	tempCard := models.Card{
		AvaliableLimit:  100,
		Currency:        "BRL",
		CVV:             123,
		DueDate:         2,
		ExpirationMonth: 2,
		ExpirationYear:  22,
		Number:          "2122123233233122",
		RealLimit:       322,
		WalletId:        2,
	}

	testDBCall(db, func(rs app.RequestScope, c routing.Context) {
		card, err := service.CreateCard(rs, &tempCard)
		assert.Nil(t, card)
		if assert.NotNil(t, err) {
			assert.Equal(t, "You're not allowed to do this.", err.Error())
		}
	})

}

func TestCreateCardWithErrorNotPassingAllRequiredAttibutes(t *testing.T) {
	db := testdata.ResetDB()
	dao := daos.NewCardDAO()
	service := NewCardService(dao)
	tempCard := models.Card{
		AvaliableLimit:  100,
		Currency:        "BRL",
		CVV:             123,
		DueDate:         2,
		ExpirationMonth: 2,
		ExpirationYear:  22,
		RealLimit:       322,
		WalletId:        1,
	}

	testDBCall(db, func(rs app.RequestScope, c routing.Context) {
		card, err := service.CreateCard(rs, &tempCard)
		assert.Nil(t, card)
		if assert.NotNil(t, err) {
			assert.Equal(t, "number: cannot be blank.", err.Error())
		}
	})

}

func TestUpdateCard(t *testing.T) {
	db := testdata.ResetDB()
	dao := daos.NewCardDAO()
	service := NewCardService(dao)
	tempCard := models.Card{
		AvaliableLimit:  100,
		Currency:        "BRL",
		CVV:             123,
		DueDate:         2,
		ExpirationMonth: 2,
		ExpirationYear:  22,
		RealLimit:       322,
		WalletId:        1,
	}

	testDBCall(db, func(rs app.RequestScope, c routing.Context) {
		card, err := service.UpdateCard(rs, 1, &tempCard)
		assert.Nil(t, err)
		if assert.NotNil(t, card) {
			assert.Equal(t, 100.0, card.AvaliableLimit)
		}
	})

}

func TestUpdateCardWithErrorWalletNotBelongToTheAuthenticatedPerson(t *testing.T) {
	db := testdata.ResetDB()
	dao := daos.NewCardDAO()
	service := NewCardService(dao)
	tempCard := models.Card{
		AvaliableLimit:  100,
		Currency:        "BRL",
		CVV:             123,
		DueDate:         2,
		ExpirationMonth: 2,
		ExpirationYear:  22,
		RealLimit:       322,
		WalletId:        2,
	}

	testDBCall(db, func(rs app.RequestScope, c routing.Context) {
		card, err := service.UpdateCard(rs, 1, &tempCard)
		assert.Nil(t, card)
		if assert.NotNil(t, err) {
			assert.Equal(t, "You're not allowed to do this.", err.Error())
		}
	})

}

func TestUpdateCardWithErrorCardNotBelongToTheAuthenticatedPerson(t *testing.T) {
	db := testdata.ResetDB()
	dao := daos.NewCardDAO()
	service := NewCardService(dao)
	tempCard := models.Card{
		AvaliableLimit:  100,
		Currency:        "BRL",
		CVV:             123,
		DueDate:         2,
		ExpirationMonth: 2,
		ExpirationYear:  22,
		RealLimit:       322,
		WalletId:        1,
	}

	testDBCall(db, func(rs app.RequestScope, c routing.Context) {
		card, err := service.UpdateCard(rs, 4, &tempCard)
		assert.Nil(t, card)
		if assert.NotNil(t, err) {
			assert.Equal(t, "You're not allowed to do this.", err.Error())
		}
	})

}

func TestDeleteCard(t *testing.T) {
	db := testdata.ResetDB()
	dao := daos.NewCardDAO()
	service := NewCardService(dao)

	testDBCall(db, func(rs app.RequestScope, c routing.Context) {
		card, err := service.DeleteCard(rs, 1)
		assert.Nil(t, err)
		if assert.NotNil(t, card) {
			assert.Equal(t, 1, card.Id)
		}
	})

}

func TestDeleteCardWithErrorCardNotBelongToTheAuthenticatedPerson(t *testing.T) {
	db := testdata.ResetDB()
	dao := daos.NewCardDAO()
	service := NewCardService(dao)

	testDBCall(db, func(rs app.RequestScope, c routing.Context) {
		card, err := service.DeleteCard(rs, 4)
		assert.Nil(t, card)
		if assert.NotNil(t, err) {
			assert.Equal(t, "You're not allowed to do this.", err.Error())
		}
	})

}
