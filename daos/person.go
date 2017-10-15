package daos

import (
	"github.com/allisonverdam/best-credit-card/app"
	"github.com/allisonverdam/best-credit-card/models"
	"github.com/go-ozzo/ozzo-dbx"
)

// PersonDAO faz a persistencia dos dados no bd
type PersonDAO struct{}

// NewPersonDAO cria um novo PersonDAO
func NewPersonDAO() *PersonDAO {
	return &PersonDAO{}
}

// Get reads the person with the specified ID from the database.
func (dao *PersonDAO) Get(rs app.RequestScope, id int) (*models.Person, error) {
	person := models.Person{}
	err := rs.Tx().Select().Model(id, &person)
	return &person, err
}

// Get reads the person with the specified ID from the database.
func (dao *PersonDAO) GetPersonByUserName(rs app.RequestScope, username string) (*models.Person, error) {
	person := models.Person{}
	err := rs.Tx().Select().Where(dbx.Like("username", username)).One(&person)
	return &person, err
}

// Create saves a new person record in the database.
// The Person.Id field will be populated with an automatically generated ID upon successful saving.
func (dao *PersonDAO) Create(rs app.RequestScope, person *models.Person) error {
	person.Id = 0
	return rs.Tx().Model(person).Insert()
}

// Update saves the changes to an person in the database.
func (dao *PersonDAO) Update(rs app.RequestScope, id int, person *models.Person) error {
	if _, err := dao.Get(rs, id); err != nil {
		return err
	}
	person.Id = id
	return rs.Tx().Model(person).Exclude("Id").Update()
}

// Delete deletes an person with the specified ID from the database.
func (dao *PersonDAO) Delete(rs app.RequestScope, id int) error {
	person, err := dao.Get(rs, id)
	if err != nil {
		return err
	}
	return rs.Tx().Model(person).Delete()
}

// Count returns the number of the person records in the database.
func (dao *PersonDAO) Count(rs app.RequestScope) (int, error) {
	var count int
	err := rs.Tx().Select("COUNT(*)").From("person").Row(&count)
	return count, err
}

// Query retrieves the person records with the specified offset and limit from the database.
func (dao *PersonDAO) Query(rs app.RequestScope, offset, limit int) ([]models.Person, error) {
	persons := []models.Person{}
	err := rs.Tx().Select().OrderBy("id").Offset(int64(offset)).Limit(int64(limit)).All(&persons)
	return persons, err
}
