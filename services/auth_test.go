package services

import (
	"testing"

	"github.com/allisonverdam/best-credit-card/app"
	"github.com/allisonverdam/best-credit-card/daos"
	"github.com/allisonverdam/best-credit-card/models"
	"github.com/allisonverdam/best-credit-card/testdata"
	"github.com/go-ozzo/ozzo-routing"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	db := testdata.ResetDB()
	dao := daos.NewPersonDAO()
	service := NewAuthService(dao)

	cred := models.Credential{Username: "allisonverdam", Password: "12345678"}

	testDBCall(db, func(rs app.RequestScope, c routing.Context) {
		token, err := service.Login(&c, cred, app.Config.JWTSigningKey)
		assert.Nil(t, err)
		assert.NotNil(t, token)
	})

}

func TestLoginWithErrorLoginOrPasswordInvalid(t *testing.T) {
	db := testdata.ResetDB()
	dao := daos.NewPersonDAO()
	service := NewAuthService(dao)

	cred := models.Credential{Username: "allisonverdam", Password: "1234"}

	testDBCall(db, func(rs app.RequestScope, c routing.Context) {
		token, err := service.Login(&c, cred, app.Config.JWTSigningKey)
		assert.Equal(t, "", token)
		if assert.NotNil(t, err) {
			assert.Equal(t, "UNAUTHORIZED", err.Error())
		}
	})

}

func TestVerifyPersonOwner(t *testing.T) {
	db := testdata.ResetDB()

	testDBCall(db, func(rs app.RequestScope, c routing.Context) {
		err := VerifyPersonOwner(rs, 1, "resource")
		assert.Nil(t, err)
	})
}

func TestVerifyPersonOwnerWithError(t *testing.T) {
	db := testdata.ResetDB()

	testDBCall(db, func(rs app.RequestScope, c routing.Context) {
		err := VerifyPersonOwner(rs, 2, "resource")
		if assert.NotNil(t, err) {
			assert.Equal(t, "FORBIDDEN", err.Error())
		}
	})
}

func TestRegister(t *testing.T) {
	db := testdata.ResetDB()
	dao := daos.NewPersonDAO()
	service := NewAuthService(dao)
	person := models.Person{Email: "joao@gmail.com", Name: "joao", Username: "joao", Password: "12345678"}

	testDBCall(db, func(rs app.RequestScope, c routing.Context) {
		rs.SetUserID(0)
		personRes, err := service.Register(rs, &person)
		assert.NotNil(t, personRes)
		assert.Nil(t, err)
	})
}

func TestRegisterWithErrorPasswordBlank(t *testing.T) {
	db := testdata.ResetDB()
	dao := daos.NewPersonDAO()
	service := NewAuthService(dao)
	person := models.Person{Email: "joao@gmail.com", Name: "joao", Username: "joao"}

	testDBCall(db, func(rs app.RequestScope, c routing.Context) {
		rs.SetUserID(0)
		personRes, err := service.Register(rs, &person)
		assert.Nil(t, personRes)
		if assert.NotNil(t, err) {
			assert.Equal(t, "password: cannot be blank.", err.Error())
		}
	})
}
