package daos

import (
	"testing"

	"github.com/allisonverdam/best-credit-card/app"
	"github.com/allisonverdam/best-credit-card/models"
	"github.com/allisonverdam/best-credit-card/testdata"
	"github.com/stretchr/testify/assert"
)

func TestGetWallet(t *testing.T) {
	db := testdata.ResetDB()
	dao := NewWalletDAO()

	testDBCall(db, func(rs app.RequestScope) {
		wallet, err := dao.GetWallet(rs, 1)
		assert.Nil(t, err)
		if assert.NotNil(t, wallet) {
			assert.Equal(t, 1, wallet.Id)
		}
	})
}

func TestGetWalletOfAuthenticatedPerson(t *testing.T) {
	db := testdata.ResetDB()
	dao := NewWalletDAO()

	testDBCall(db, func(rs app.RequestScope) {
		wallets, err := dao.GetAuthenticatedPersonWallets(rs, rs.UserID())
		assert.Nil(t, err)
		if assert.NotNil(t, wallets) {
			assert.Equal(t, 1, wallets[0].Id)
		}
	})
}

func TestCreateWallet(t *testing.T) {
	db := testdata.ResetDB()
	dao := NewWalletDAO()

	testDBCall(db, func(rs app.RequestScope) {
		wallet := &models.Wallet{
			MaximumLimit: 500,
			CurrentLimit: 200,
			PersonId:     1,
		}
		err := dao.CreateWallet(rs, wallet)
		assert.Nil(t, err)
		assert.NotZero(t, wallet.Id)
	})
}

func TestCreateWalletWithErrorPersonNull(t *testing.T) {
	db := testdata.ResetDB()
	dao := NewWalletDAO()

	testDBCall(db, func(rs app.RequestScope) {
		wallet := &models.Wallet{
			MaximumLimit: 500,
			CurrentLimit: 200,
		}
		err := dao.CreateWallet(rs, wallet)
		if assert.NotNil(t, err) {
			assert.Equal(t, "person_id: cannot be blank.", err.Error())
		}
	})
}

func TestUpdateWallet(t *testing.T) {
	db := testdata.ResetDB()
	dao := NewWalletDAO()

	testDBCall(db, func(rs app.RequestScope) {
		wallet := &models.Wallet{
			Id:           1,
			CurrentLimit: 234,
			MaximumLimit: 200,
			PersonId:     1,
		}
		err := dao.UpdateWallet(rs, wallet.Id, wallet)
		assert.Nil(t, err)
	})
}

func TestUpdateWalletWithErrorNotExistWalletWithThisID(t *testing.T) {
	db := testdata.ResetDB()
	dao := NewWalletDAO()

	testDBCall(db, func(rs app.RequestScope) {
		wallet := &models.Wallet{
			MaximumLimit: 42,
			CurrentLimit: 22,
			PersonId:     1,
		}
		err := dao.UpdateWallet(rs, 99, wallet)
		if assert.NotNil(t, err) {
			assert.Equal(t, "sql: no rows in result set", err.Error())
		}
	})
}

func TestUpdateWalletWithErrorPersonIdISNull(t *testing.T) {
	db := testdata.ResetDB()
	dao := NewWalletDAO()

	testDBCall(db, func(rs app.RequestScope) {
		wallet := &models.Wallet{
			Id:           0,
			MaximumLimit: 42,
		}
		err := dao.UpdateWallet(rs, wallet.Id, wallet)
		assert.NotNil(t, err)
		assert.Equal(t, "person_id: cannot be blank.", err.Error())
	})
}

func TestDeleteWallet(t *testing.T) {
	db := testdata.ResetDB()
	dao := NewWalletDAO()

	testDBCall(db, func(rs app.RequestScope) {
		err := dao.DeleteWallet(rs, 1)
		assert.Nil(t, err)
	})
}

func TestDeleteWalletWithErrorWalletNotFound(t *testing.T) {
	db := testdata.ResetDB()
	dao := NewWalletDAO()

	testDBCall(db, func(rs app.RequestScope) {
		err := dao.DeleteWallet(rs, 99999)
		assert.NotNil(t, err)
	})
}
