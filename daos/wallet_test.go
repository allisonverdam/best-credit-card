package daos

import (
	"testing"

	"github.com/allisonverdam/best-credit-card/app"
	"github.com/allisonverdam/best-credit-card/models"
	"github.com/allisonverdam/best-credit-card/testdata"
	"github.com/stretchr/testify/assert"
)

func TestWalletDAO(t *testing.T) {
	db := testdata.ResetDB()
	dao := NewWalletDAO()

	{
		// Get
		testDBCall(db, func(rs app.RequestScope) {
			wallet, err := dao.Get(rs, 2)
			assert.Nil(t, err)
			if assert.NotNil(t, wallet) {
				assert.Equal(t, 2, wallet.Id)
			}
		})
	}

	{
		// Create
		testDBCall(db, func(rs app.RequestScope) {
			wallet := &models.Wallet{
				MaximumLimit: 500,
				PersonId:     1,
				RealLimit:    200,
			}
			err := dao.Create(rs, wallet)
			assert.Nil(t, err)
			assert.NotZero(t, wallet.Id)
		})
	}

	{
		// Create with error
		testDBCall(db, func(rs app.RequestScope) {
			wallet := &models.Wallet{
				Id:           1000,
				MaximumLimit: 700,
				PersonId:     99999999,
			}
			err := dao.Create(rs, wallet)
			assert.NotNil(t, err)
			assert.Equal(t, "real_limit: cannot be blank.", err.Error())
		})
	}

	{
		// Create with error
		testDBCall(db, func(rs app.RequestScope) {
			wallet := &models.Wallet{
				Id:           1000,
				MaximumLimit: 700,
				PersonId:     99999999,
				RealLimit:    100,
			}
			err := dao.Create(rs, wallet)
			assert.NotNil(t, err)
			assert.Equal(t, "pq: insert or update on table \"wallet\" violates foreign key constraint \"wallet_person_id_fkey\"", err.Error())
		})
	}

	{
		// Update
		testDBCall(db, func(rs app.RequestScope) {
			wallet := &models.Wallet{
				Id:           1,
				RealLimit:    234,
				PersonId:     1,
				MaximumLimit: 200,
			}
			err := dao.Update(rs, wallet.Id, wallet)
			assert.Nil(t, err)
		})
	}

	{
		// Update with error
		testDBCall(db, func(rs app.RequestScope) {
			wallet := &models.Wallet{
				Id:           0,
				MaximumLimit: 42,
				PersonId:     1,
				RealLimit:    22,
			}
			err := dao.Update(rs, wallet.Id, wallet)
			assert.NotNil(t, err)
			assert.Equal(t, "sql: no rows in result set", err.Error())
		})
	}

	{
		// Update with error
		testDBCall(db, func(rs app.RequestScope) {
			wallet := &models.Wallet{
				Id:           0,
				MaximumLimit: 42,
				PersonId:     1,
			}
			err := dao.Update(rs, wallet.Id, wallet)
			assert.NotNil(t, err)
			assert.Equal(t, "real_limit: cannot be blank.", err.Error())
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
