package daos

import (
	"testing"

	"github.com/allisonverdam/best-credit-card/app"
	"github.com/allisonverdam/best-credit-card/models"
	"github.com/allisonverdam/best-credit-card/testdata"
	"github.com/stretchr/testify/assert"
)

func TestGetWalletCardsLimits(t *testing.T) {
	db := testdata.ResetDB()
	dao := NewCardDAO()

	testDBCall(db, func(rs app.RequestScope) {
		card, err := dao.GetWalletCardsLimits(rs, 2)
		assert.Nil(t, err)
		if assert.NotNil(t, card) {
			assert.Equal(t, 1430.0, card.AvaliableLimit)
			assert.Equal(t, 1800.0, card.RealLimit)
		}
	})
}

func TestGetCard(t *testing.T) {
	db := testdata.ResetDB()
	dao := NewCardDAO()

	testDBCall(db, func(rs app.RequestScope) {
		card, err := dao.GetCard(rs, 2)
		assert.Nil(t, err)
		if assert.NotNil(t, card) {
			assert.Equal(t, 2, card.Id)
		}
	})
}

func TestGetCardsByWallet(t *testing.T) {
	db := testdata.ResetDB()
	dao := NewCardDAO()

	testDBCall(db, func(rs app.RequestScope) {
		wallet := models.Wallet{}
		wallet.PersonId = 1
		wallet.Id = 1

		cards, err := dao.GetCardsByWallet(rs, wallet)
		assert.Nil(t, err)
		if assert.NotNil(t, &cards) {
			assert.Equal(t, 4, len(*&cards))
		}
	})
}

func TestGetBestCardsByWallet(t *testing.T) {
	db := testdata.ResetDB()
	dao := NewCardDAO()

	testDBCall(db, func(rs app.RequestScope) {
		wallet := models.Wallet{}
		wallet.PersonId = 1
		wallet.Id = 1

		cards, err := dao.GetBestCardsByWallet(rs, wallet)
		assert.Nil(t, err)
		if assert.NotNil(t, cards) {
			assert.Equal(t, 4, len(cards))
		}
	})
}

func TestCreateCard(t *testing.T) {
	db := testdata.ResetDB()
	dao := NewCardDAO()

	testDBCall(db, func(rs app.RequestScope) {
		card := &models.Card{
			Id:              1000,
			Number:          "1234123412341299",
			RealLimit:       100,
			AvaliableLimit:  50,
			CVV:             123,
			DueDate:         22,
			ExpirationMonth: 01,
			Currency:        "BRL",
			ExpirationYear:  17,
			WalletId:        1,
		}
		err := dao.CreateCard(rs, card)
		assert.Nil(t, err)
		assert.Equal(t, 1000, card.Id)
		assert.NotZero(t, card.Id)
	})
}

func TestUpdateCard(t *testing.T) {
	db := testdata.ResetDB()
	dao := NewCardDAO()

	testDBCall(db, func(rs app.RequestScope) {
		card := &models.Card{
			Id:       2,
			Number:   "1234123412341298",
			WalletId: 1,
		}
		err := dao.UpdateCard(rs, card.Id, card)
		assert.Nil(t, err)
	})
}

func TestUpdateCardWithErrorCardNotFound(t *testing.T) {
	db := testdata.ResetDB()
	dao := NewCardDAO()

	testDBCall(db, func(rs app.RequestScope) {
		card := &models.Card{
			Id:       0,
			Number:   "1234123412341297",
			WalletId: 1,
		}
		err := dao.UpdateCard(rs, card.Id, card)
		if assert.NotNil(t, err) {
			assert.Equal(t, "sql: no rows in result set", err.Error())
		}
	})
}

func TestDeleteCard(t *testing.T) {
	db := testdata.ResetDB()
	dao := NewCardDAO()

	testDBCall(db, func(rs app.RequestScope) {
		err := dao.DeleteCard(rs, 2)
		assert.Nil(t, err)
	})
}

func TestDeleteCardWithErrorCardNotFound(t *testing.T) {
	db := testdata.ResetDB()
	dao := NewCardDAO()

	testDBCall(db, func(rs app.RequestScope) {
		err := dao.DeleteCard(rs, 99999)
		if assert.NotNil(t, err) {
			assert.Equal(t, "sql: no rows in result set", err.Error())
		}
	})
}
