package daos

import (
	"testing"

	"github.com/allisonverdam/best-credit-card/app"
	"github.com/allisonverdam/best-credit-card/models"
	"github.com/allisonverdam/best-credit-card/testdata"
	"github.com/stretchr/testify/assert"
)

func TestGetPerson(t *testing.T) {
	db := testdata.ResetDB()
	dao := NewPersonDAO()

	testDBCall(db, func(rs app.RequestScope) {
		person, err := dao.GetPerson(rs, 2)
		assert.Nil(t, err)
		if assert.NotNil(t, person) {
			assert.Equal(t, 2, person.Id)
		}
	})

}

func TestGetPersonWithoutPassword(t *testing.T) {
	db := testdata.ResetDB()
	dao := NewPersonDAO()

	testDBCall(db, func(rs app.RequestScope) {
		person, err := dao.GetPersonWithoutPassword(rs, 2)
		assert.Nil(t, err)
		if assert.NotNil(t, person) {
			assert.Equal(t, "", person.Password)
		}
	})

}

func TestGetPersonWithoutPasswordWithErrorPersonNotFound(t *testing.T) {
	db := testdata.ResetDB()
	dao := NewPersonDAO()

	testDBCall(db, func(rs app.RequestScope) {
		person, err := dao.GetPersonWithoutPassword(rs, 9999)
		assert.Nil(t, person)
		if assert.NotNil(t, err) {
			assert.Equal(t, "sql: no rows in result set", err.Error())
		}
	})

}

func TestGetPersonByUserName(t *testing.T) {
	db := testdata.ResetDB()
	dao := NewPersonDAO()

	testDBCall(db, func(rs app.RequestScope) {
		person, err := dao.GetPersonByUserName(rs, "allisonverdam")
		assert.Nil(t, err)
		if assert.NotNil(t, person) {
			assert.Equal(t, 1, person.Id)
		}
	})

}

func TestGetPersonByUserNameWithErrorUserNameNotFound(t *testing.T) {
	db := testdata.ResetDB()
	dao := NewPersonDAO()

	testDBCall(db, func(rs app.RequestScope) {
		person, err := dao.GetPersonByUserName(rs, "xaxax2112")
		assert.Nil(t, person)
		if assert.NotNil(t, err) {
			assert.Equal(t, "sql: no rows in result set", err.Error())
		}
	})

}

func TestCreatePerson(t *testing.T) {
	db := testdata.ResetDB()
	dao := NewPersonDAO()

	testDBCall(db, func(rs app.RequestScope) {
		person := &models.Person{
			Id:       1000,
			Email:    "josh@gmail.com",
			Name:     "Josh",
			Password: "12345678",
			Username: "josh123",
		}
		err := dao.CreatePerson(rs, person)
		assert.Nil(t, err)
		assert.Equal(t, 1000, person.Id)
		assert.NotZero(t, person.Id)
	})

}

func TestUpdatePerson(t *testing.T) {
	db := testdata.ResetDB()
	dao := NewPersonDAO()

	testDBCall(db, func(rs app.RequestScope) {
		person := &models.Person{
			Id:       2,
			Name:     "Carlos",
			Email:    "carlos@g.com",
			Password: "12345678",
			Username: "carlos",
		}
		err := dao.UpdatePerson(rs, person.Id, person)
		assert.Nil(t, err)
	})

}

func TestUpdatePersonWithErrorPersonNotFound(t *testing.T) {
	db := testdata.ResetDB()
	dao := NewPersonDAO()

	testDBCall(db, func(rs app.RequestScope) {
		person := &models.Person{
			Id:       99999,
			Email:    "email.com",
			Password: "12345678",
		}
		err := dao.UpdatePerson(rs, person.Id, person)
		if assert.NotNil(t, err) {
			assert.Equal(t, "sql: no rows in result set", err.Error())
		}
	})

}

func TestUpdatePersonWithErrorUserNameDuplicated(t *testing.T) {
	db := testdata.ResetDB()
	dao := NewPersonDAO()

	testDBCall(db, func(rs app.RequestScope) {
		person := &models.Person{
			Id:       3,
			Name:     "Ana",
			Email:    "ana@g.com",
			Username: "allisonverdam",
		}
		err := dao.UpdatePerson(rs, person.Id, person)
		if assert.NotNil(t, err) {
			assert.Equal(t, "pq: duplicate key value violates unique constraint \"person_username_key\"", err.Error())
		}
	})

}

func TestDeletePerson(t *testing.T) {
	db := testdata.ResetDB()
	dao := NewPersonDAO()

	testDBCall(db, func(rs app.RequestScope) {
		err := dao.DeletePerson(rs, 2)
		assert.Nil(t, err)
	})

}

func TestDeletePersonWithErrorPersonNotFound(t *testing.T) {
	db := testdata.ResetDB()
	dao := NewPersonDAO()

	testDBCall(db, func(rs app.RequestScope) {
		err := dao.DeletePerson(rs, 99999)
		assert.NotNil(t, err)
	})

}
