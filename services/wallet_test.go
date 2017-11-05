package services

import (
	"testing"

	"github.com/allisonverdam/best-credit-card/daos"
	"github.com/allisonverdam/best-credit-card/models"

	"github.com/allisonverdam/best-credit-card/app"
	"github.com/allisonverdam/best-credit-card/testdata"
	routing "github.com/go-ozzo/ozzo-routing"
	"github.com/stretchr/testify/assert"
)

func TestGetAuthenticatedPersonWallets(t *testing.T) {
	db := testdata.ResetDB()

	testDBCall(db, func(rs app.RequestScope, c routing.Context) {
		wallets, err := GetAuthenticatedPersonWallets(rs)
		assert.Nil(t, err)
		if assert.NotNil(t, wallets) {
			assert.Equal(t, 1, len(wallets))
		}
	})

}

func TestCreateWallet(t *testing.T) {
	dao := daos.NewWalletDAO()
	service := NewWalletService(dao)
	db := testdata.ResetDB()
	wallet := models.Wallet{
		AvaliableLimit: 125,
		CurrentLimit:   75,
		MaximumLimit:   200,
		PersonId:       1,
	}

	testDBCall(db, func(rs app.RequestScope, c routing.Context) {
		wallet, err := service.CreateWallet(rs, &wallet)
		assert.Nil(t, err)
		if assert.NotNil(t, wallet) {
			assert.Equal(t, 125.0, wallet.AvaliableLimit)
		}
	})

}

func TestUpdateWallet(t *testing.T) {
	dao := daos.NewWalletDAO()
	service := NewWalletService(dao)
	db := testdata.ResetDB()
	wallet := models.Wallet{
		CurrentLimit: 0,
		PersonId:     1,
	}

	testDBCall(db, func(rs app.RequestScope, c routing.Context) {
		wallet, err := service.UpdateWallet(rs, 1, &wallet)
		assert.Nil(t, err)
		if assert.NotNil(t, wallet) {
			assert.Equal(t, 0.0, wallet.MaximumLimit)
		}
	})

}

func TestUpdateWalletWithError(t *testing.T) {
	dao := daos.NewWalletDAO()
	service := NewWalletService(dao)
	db := testdata.ResetDB()
	wallet := models.Wallet{
		CurrentLimit: 1000,
		PersonId:     1,
	}

	testDBCall(db, func(rs app.RequestScope, c routing.Context) {
		wallet, err := service.UpdateWallet(rs, 1, &wallet)
		assert.Nil(t, wallet)
		if assert.NotNil(t, err) {
			assert.Equal(t, "Attempted to increase to a higher limit than available. This is your avaliable limit: R$0. Pay some credit card to release more credit.", err.Error())
		}
	})

}

func TestDeleteWallet(t *testing.T) {
	dao := daos.NewWalletDAO()
	service := NewWalletService(dao)
	db := testdata.ResetDB()

	testDBCall(db, func(rs app.RequestScope, c routing.Context) {
		wallet, err := service.DeleteWallet(rs, 1)
		assert.Nil(t, err)
		if assert.NotNil(t, wallet) {
			assert.Equal(t, 1, wallet.Id)
		}
	})

}

func TestDeleteWalletWithError(t *testing.T) {
	dao := daos.NewWalletDAO()
	service := NewWalletService(dao)
	db := testdata.ResetDB()

	testDBCall(db, func(rs app.RequestScope, c routing.Context) {
		wallet, err := service.DeleteWallet(rs, 2)
		assert.Nil(t, wallet)
		if assert.NotNil(t, err) {
			assert.Equal(t, "You're not allowed to do this.", err.Error())
		}
	})

}
