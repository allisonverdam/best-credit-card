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

func TestGetAuthenticatedPersonWallet(t *testing.T) {
	db := testdata.ResetDB()
	dao := daos.NewWalletDAO()
	service := NewWalletService(dao)

	testDBCall(db, func(rs app.RequestScope, c routing.Context) {
		wallet, err := service.GetAuthenticatedPersonWallet(rs)
		assert.Nil(t, err)
		if assert.NotNil(t, wallet) {
			assert.Equal(t, 1, wallet.Id)
		}
	})

}

func TestCreateWallet(t *testing.T) {
	dao := daos.NewWalletDAO()
	service := NewWalletService(dao)
	db := testdata.ResetDB()
	wallet := models.Wallet{
		PersonId: 1,
	}

	testDBCall(db, func(rs app.RequestScope, c routing.Context) {
		wallet, err := service.CreateWallet(rs, &wallet)
		assert.Nil(t, err)
		if assert.NotNil(t, wallet) {
			assert.Equal(t, 4, wallet.Id)
		}
	})

}

func TestCreateWalletWithErrorWalletNotBelongToTheAuthenticatedPerson(t *testing.T) {
	dao := daos.NewWalletDAO()
	service := NewWalletService(dao)
	db := testdata.ResetDB()
	wallet := models.Wallet{
		PersonId: 21,
	}

	testDBCall(db, func(rs app.RequestScope, c routing.Context) {
		wallet, err := service.CreateWallet(rs, &wallet)
		assert.Nil(t, wallet)
		if assert.NotNil(t, err) {
			assert.Equal(t, "You're not allowed to do this.", err.Error())
		}
	})

}

func TestCreateWalletWithErrorNotPassingAllRequiredParams(t *testing.T) {
	dao := daos.NewWalletDAO()
	service := NewWalletService(dao)
	db := testdata.ResetDB()
	wallet := models.Wallet{}

	testDBCall(db, func(rs app.RequestScope, c routing.Context) {
		wallet, err := service.CreateWallet(rs, &wallet)
		assert.Nil(t, wallet)
		if assert.NotNil(t, err) {
			assert.Equal(t, "-: cannot be blank.", err.Error())
		}
	})

}

func TestUpdateAuthenticatedPersonWallet(t *testing.T) {
	dao := daos.NewWalletDAO()
	service := NewWalletService(dao)
	db := testdata.ResetDB()
	wallet := models.Wallet{
		CurrentLimit: 0,
		PersonId:     1,
	}

	testDBCall(db, func(rs app.RequestScope, c routing.Context) {
		wallet, err := service.UpdateAuthenticatedPersonWallet(rs, &wallet)
		assert.Nil(t, err)
		if assert.NotNil(t, wallet) {
			assert.Equal(t, 1550.0, wallet.MaximumLimit)
		}
	})

}

func TestUpdateAuthenticatedPersonWalletWithErrorInvalidCurrentLimit(t *testing.T) {
	dao := daos.NewWalletDAO()
	service := NewWalletService(dao)
	db := testdata.ResetDB()
	wallet := models.Wallet{
		CurrentLimit: 2000,
		PersonId:     1,
	}

	testDBCall(db, func(rs app.RequestScope, c routing.Context) {
		wallet, err := service.UpdateAuthenticatedPersonWallet(rs, &wallet)
		assert.Nil(t, wallet)
		if assert.NotNil(t, err) {
			assert.Equal(t, "Attempted to increase to a higher limit than available. This is your avaliable limit: R$1030. Pay some credit card to release more credit.", err.Error())
		}
	})

}

func TestUpdateAuthenticatedPersonWalletWithErrorUserIdNull(t *testing.T) {
	dao := daos.NewWalletDAO()
	service := NewWalletService(dao)
	db := testdata.ResetDB()
	wallet := models.Wallet{
		CurrentLimit: 10,
		PersonId:     1,
	}

	testDBCall(db, func(rs app.RequestScope, c routing.Context) {
		rs.SetUserID(0)
		wallet, err := service.UpdateAuthenticatedPersonWallet(rs, &wallet)
		assert.Nil(t, wallet)
		if assert.NotNil(t, err) {
			assert.Equal(t, "sql: no rows in result set", err.Error())
		}
	})

}

func TestGetWalletThrowVerification(t *testing.T) {
	dao := daos.NewWalletDAO()
	service := NewWalletService(dao)
	db := testdata.ResetDB()

	testDBCall(db, func(rs app.RequestScope, c routing.Context) {
		wallet, err := service.GetWalletThrowVerification(rs, 1)
		assert.Nil(t, err)
		if assert.NotNil(t, wallet) {
			assert.Equal(t, 1, wallet.Id)
		}
	})

}

func TestGetWalletThrowVerificationWithError(t *testing.T) {
	dao := daos.NewWalletDAO()
	service := NewWalletService(dao)
	db := testdata.ResetDB()

	testDBCall(db, func(rs app.RequestScope, c routing.Context) {
		wallet, err := service.GetWalletThrowVerification(rs, 999)
		assert.Nil(t, wallet)
		if assert.NotNil(t, err) {
			assert.Equal(t, "sql: no rows in result set", err.Error())
		}
	})

}

func TestUpdateWalletLimitsWithErrorWalletNotFound(t *testing.T) {
	dao := daos.NewWalletDAO()
	service := NewWalletService(dao)
	db := testdata.ResetDB()
	card := models.Card{
		WalletId: 999,
	}

	testDBCall(db, func(rs app.RequestScope, c routing.Context) {
		err := service.UpdateWalletLimits(rs, card)
		if assert.NotNil(t, err) {
			assert.Equal(t, "sql: no rows in result set", err.Error())
		}
	})

}

func TestUpdateWalletLimitsWithErrorWalletNotBelongToTheAuthenticatedPerson(t *testing.T) {
	dao := daos.NewWalletDAO()
	service := NewWalletService(dao)
	db := testdata.ResetDB()
	card := models.Card{
		WalletId: 2,
	}

	testDBCall(db, func(rs app.RequestScope, c routing.Context) {
		err := service.UpdateWalletLimits(rs, card)
		if assert.NotNil(t, err) {
			assert.Equal(t, "You're not allowed to do this.", err.Error())
		}
	})

}
