package daos

import (
	"testing"

	"github.com/allisonverdam/best-credit-card/app"
	"github.com/allisonverdam/best-credit-card/models"
	"github.com/allisonverdam/best-credit-card/testdata"
	"github.com/stretchr/testify/assert"
)

func TestCardDAO(t *testing.T) {
	db := testdata.ResetDB()
	dao := NewCardDAO()

	{
		// Get
		testDBCall(db, func(rs app.RequestScope) {
			card, err := dao.Get(rs, 2)
			assert.Nil(t, err)
			if assert.NotNil(t, card) {
				assert.Equal(t, 2, card.Id)
			}
		})
	}

	{
		// GetBestCardsByWalletId
		testDBCall(db, func(rs app.RequestScope) {
			personId := 1
			walletId := 1
			cards, err := dao.GetBestCardsByWalletId(rs, personId, walletId)
			assert.Nil(t, err)
			if assert.NotNil(t, cards) {
				assert.Equal(t, 3, len(cards))
			}
		})
	}

	{
		// GetBestCardsByWalletId with error
		//wallet not belong to person
		testDBCall(db, func(rs app.RequestScope) {
			personId := 1
			walletId := 2
			card, err := dao.GetBestCardsByWalletId(rs, personId, walletId)
			assert.Nil(t, card)
			assert.Equal(t, "FORBIDDEN", err.Error())
		})
	}

	{
		// GetCardsByWalletId with error
		//wallet not exist
		testDBCall(db, func(rs app.RequestScope) {
			personId := 1
			walletId := 0
			card, err := dao.GetBestCardsByWalletId(rs, personId, walletId)
			assert.Nil(t, card)
			assert.Equal(t, "sql: no rows in result set", err.Error())
		})
	}

	{
		// GetCardsByWalletId
		testDBCall(db, func(rs app.RequestScope) {
			personId := 1
			walletId := 1
			cards, err := dao.GetCardsByWalletId(rs, personId, walletId)
			assert.Nil(t, err)
			if assert.NotNil(t, cards) {
				assert.Equal(t, 3, len(cards))
			}
		})
	}

	{
		// GetCardsByWalletId with error
		//wallet not belong to person
		testDBCall(db, func(rs app.RequestScope) {
			personId := 1
			walletId := 2
			card, err := dao.GetCardsByWalletId(rs, personId, walletId)
			assert.Nil(t, card)
			assert.Equal(t, "FORBIDDEN", err.Error())
		})
	}

	{
		// GetCardsByWalletId with error
		//wallet not exist
		testDBCall(db, func(rs app.RequestScope) {
			personId := 1
			walletId := 0
			card, err := dao.GetCardsByWalletId(rs, personId, walletId)
			assert.Nil(t, card)
			assert.Equal(t, "sql: no rows in result set", err.Error())
		})
	}

	{
		// Create
		testDBCall(db, func(rs app.RequestScope) {
			card := &models.Card{
				Id:              1000,
				Number:          "1234123412341299",
				RealLimit:       100,
				AvaliableLimit:  50,
				CVV:             123,
				DueDate:         22,
				ExpirationMonth: 01,
				ExpirationYear:  17,
				WalletId:        1,
			}
			err := dao.Create(rs, card)
			assert.Nil(t, err)
			assert.Equal(t, 1000, card.Id)
			assert.NotZero(t, card.Id)
		})
	}

	{
		// Update
		testDBCall(db, func(rs app.RequestScope) {
			card := &models.Card{
				Id:       2,
				Number:   "1234123412341298",
				WalletId: 1,
			}
			err := dao.Update(rs, card.Id, card)
			assert.Nil(t, err)
		})
	}

	{
		// Update with error
		testDBCall(db, func(rs app.RequestScope) {
			card := &models.Card{
				Id:       2,
				Number:   "1234123412341297",
				WalletId: 1,
			}
			err := dao.Update(rs, 0, card)
			assert.NotNil(t, err)
		})
	}

	{
		// Delete
		testDBCall(db, func(rs app.RequestScope) {
			err := dao.Delete(rs, 2)
			assert.Nil(t, err)
		})
	}

	{
		// Delete with error
		testDBCall(db, func(rs app.RequestScope) {
			err := dao.Delete(rs, 99999)
			assert.NotNil(t, err)
		})
	}
}
