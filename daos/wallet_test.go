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
			wallet, err := dao.Get(rs, 1)
			assert.Nil(t, err)
			if assert.NotNil(t, wallet) {
				assert.Equal(t, 1, wallet.Id)
			}
		})
	}

	{
		// Get
		testDBCall(db, func(rs app.RequestScope) {
			wallets, err := dao.GetAuthenticatedPersonWallets(rs, rs.UserID())
			assert.Nil(t, err)
			if assert.NotNil(t, wallets) {
				assert.Equal(t, 1, wallets[0].Id)
			}
		})
	}

	{
		// Create
		testDBCall(db, func(rs app.RequestScope) {
			wallet := &models.Wallet{
				MaximumLimit: 500,
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
				RealLimit:    0,
			}
			err := dao.Create(rs, wallet)
			assert.NotNil(t, err)
			assert.Equal(t, "real_limit: cannot be blank.", err.Error())
		})
	}

	{
		// Update
		testDBCall(db, func(rs app.RequestScope) {
			wallet := &models.Wallet{
				Id:           1,
				RealLimit:    234,
				MaximumLimit: 200,
			}
			err := dao.Update(rs, wallet.Id, wallet)
			assert.Nil(t, err)
		})
	}

	{
		// Update with error
		//Wallet don't belong to authenticated user
		testDBCall(db, func(rs app.RequestScope) {
			wallet := &models.Wallet{
				Id:           2,
				MaximumLimit: 42,
				RealLimit:    22,
			}
			err := dao.Update(rs, wallet.Id, wallet)
			assert.NotNil(t, err)
			assert.Equal(t, "FORBIDDEN", err.Error())
		})
	}

	{
		// Update with error
		testDBCall(db, func(rs app.RequestScope) {
			wallet := &models.Wallet{
				Id:           0,
				MaximumLimit: 42,
			}
			err := dao.Update(rs, wallet.Id, wallet)
			assert.NotNil(t, err)
			assert.Equal(t, "real_limit: cannot be blank.", err.Error())
		})
	}

	{
		// Delete
		testDBCall(db, func(rs app.RequestScope) {
			err := dao.Delete(rs, 1)
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
