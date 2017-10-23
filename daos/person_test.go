package daos

import (
	"testing"

	"github.com/allisonverdam/best-credit-card/app"
	"github.com/allisonverdam/best-credit-card/models"
	"github.com/allisonverdam/best-credit-card/testdata"
	"github.com/stretchr/testify/assert"
)

func TestPersonDAO(t *testing.T) {
	db := testdata.ResetDB()
	dao := NewPersonDAO()

	{
		// Get
		testDBCall(db, func(rs app.RequestScope) {
			person, err := dao.Get(rs, 2)
			assert.Nil(t, err)
			if assert.NotNil(t, person) {
				assert.Equal(t, 2, person.Id)
			}
		})
	}

	{
		// GetWithoutPassword
		testDBCall(db, func(rs app.RequestScope) {
			person, err := dao.GetWithoutPassword(rs, 2)
			assert.Nil(t, err)
			if assert.NotNil(t, person) {
				assert.Equal(t, "", person.Password)
			}
		})
	}

	{
		// GetPersonByUserName
		testDBCall(db, func(rs app.RequestScope) {
			person, err := dao.GetPersonByUserName(rs, "allisonverdam")
			assert.Nil(t, err)
			if assert.NotNil(t, person) {
				assert.Equal(t, 1, person.Id)
			}
		})
	}

	{
		// GetPersonByUserName with error
		testDBCall(db, func(rs app.RequestScope) {
			person, err := dao.GetPersonByUserName(rs, "xaxax2112")
			assert.Nil(t, person)
			if assert.NotNil(t, err) {
				assert.Equal(t, "sql: no rows in result set", err.Error())
			}
		})
	}

	{
		// Create
		testDBCall(db, func(rs app.RequestScope) {
			person := &models.Person{
				Id:       1000,
				Email:    "josh@gmail.com",
				Name:     "Josh",
				Password: "12345678",
				Username: "josh123",
			}
			err := dao.Create(rs, person)
			assert.Nil(t, err)
			assert.Equal(t, 1000, person.Id)
			assert.NotZero(t, person.Id)
		})
	}

	{
		// Create with error
		testDBCall(db, func(rs app.RequestScope) {
			person := &models.Person{
				Id:       1000,
				Name:     "Josh",
				Password: "12345678",
				Username: "josh123",
			}
			err := dao.Create(rs, person)
			assert.NotNil(t, err)
			assert.Equal(t, "email: cannot be blank.", err.Error())
		})
	}

	{
		// Update
		testDBCall(db, func(rs app.RequestScope) {
			person := &models.Person{
				Id:       2,
				Name:     "Carlos",
				Email:    "carlos@g.com",
				Password: "12345678",
				Username: "carlos",
			}
			err := dao.Update(rs, person.Id, person)
			assert.Nil(t, err)
		})
	}

	{
		// Update with error
		testDBCall(db, func(rs app.RequestScope) {
			person := &models.Person{
				Email:    "email.com",
				Password: "12345678",
			}
			err := dao.Update(rs, 99999, person)
			assert.NotNil(t, err)
		})
	}

	{
		// Update with error
		testDBCall(db, func(rs app.RequestScope) {
			person := &models.Person{
				Email:    "ana@g.com",
				Name:     "Ana",
				Password: "12345678",
				Username: "ana123",
			}
			err := dao.Update(rs, 99999, person)
			assert.NotNil(t, err)
		})
	}

	{
		// Update with error
		// not passign password
		testDBCall(db, func(rs app.RequestScope) {
			person := &models.Person{
				Email:    "ana@g.com",
				Username: "ana123",
				Password: "",
			}
			err := dao.Update(rs, 1, person)
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
