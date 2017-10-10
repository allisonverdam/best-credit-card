package daos

import (
	"testing"
	"time"

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
		// Create
		testDBCall(db, func(rs app.RequestScope) {
			card := &models.Card{
				Id:             1000,
				Number:         "1234123412341299",
				Limit:          100,
				CVV:            123,
				DueDate:        time.Now(),
				ExpirationDate: time.Date(2022, 10, 10, 0, 0, 0, 0, time.Local),
				PersonId:       1,
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
				PersonId: 1,
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
				PersonId: 1,
			}
			err := dao.Update(rs, 99999, card)
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

	{
		// Query
		testDBCall(db, func(rs app.RequestScope) {
			cards, err := dao.Query(rs, 1, 3)
			assert.Nil(t, err)
			assert.Equal(t, 3, len(cards))
		})
	}

	{
		// Count
		testDBCall(db, func(rs app.RequestScope) {
			count, err := dao.Count(rs)
			assert.Nil(t, err)
			assert.NotZero(t, count)
		})
	}
}
