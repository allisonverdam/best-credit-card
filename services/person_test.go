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

func TestGetPerson(t *testing.T) {
	db := testdata.ResetDB()
	dao := daos.NewPersonDAO()
	service := NewPersonService(dao)

	testDBCall(db, func(rs app.RequestScope, c routing.Context) {
		person, err := service.GetPerson(rs, 1)
		assert.Nil(t, err)
		if assert.NotNil(t, person) {
			assert.Equal(t, "Allison V.", person.Name)
		}
	})

}

func TestGetPersonWithError(t *testing.T) {
	db := testdata.ResetDB()
	dao := daos.NewPersonDAO()
	service := NewPersonService(dao)

	testDBCall(db, func(rs app.RequestScope, c routing.Context) {
		person, err := service.GetPerson(rs, 2)
		assert.Nil(t, person)
		if assert.NotNil(t, err) {
			assert.Equal(t, "You're not allowed to do this.", err.Error())
		}
	})

}

func TestUpdateAuthenticatedPerson(t *testing.T) {
	db := testdata.ResetDB()
	dao := daos.NewPersonDAO()
	service := NewPersonService(dao)
	newPerson := models.Person{
		Email: "teste@g.com",
		Name:  "carlos",
	}

	testDBCall(db, func(rs app.RequestScope, c routing.Context) {
		person, err := service.UpdateAuthenticatedPerson(rs, &newPerson)
		assert.Nil(t, err)
		if assert.NotNil(t, person) {
			assert.Equal(t, "carlos", person.Name)
		}
	})

}
