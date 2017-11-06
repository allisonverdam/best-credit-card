package daos

import (
	"github.com/allisonverdam/best-credit-card/app"
	"github.com/allisonverdam/best-credit-card/models"
	"github.com/go-ozzo/ozzo-dbx"
	"golang.org/x/crypto/bcrypt"
)

// PersonDAO faz a persistencia dos dados no bd
type PersonDAO struct{}

// NewPersonDAO cria um novo PersonDAO
func NewPersonDAO() *PersonDAO {
	return &PersonDAO{}
}

func (dao *PersonDAO) GetPerson(rs app.RequestScope, id int) (*models.Person, error) {
	person := models.Person{}
	err := rs.Tx().Select().Model(id, &person)

	return &person, err
}

func (dao *PersonDAO) GetPersonWithoutPassword(rs app.RequestScope, id int) (*models.Person, error) {
	person := models.Person{}
	err := rs.Tx().Select().Model(id, &person)
	if err != nil {
		return nil, err
	}

	person.Password = ""

	return &person, err
}

func (dao *PersonDAO) GetPersonByUserName(rs app.RequestScope, username string) (*models.Person, error) {
	person := models.Person{}
	err := rs.Tx().Select().Where(dbx.HashExp{"username": username}).One(&person)
	if err != nil {
		return nil, err
	}
	return &person, nil
}

func (dao *PersonDAO) CreatePerson(rs app.RequestScope, person *models.Person) error {
	person.Password = EncryptPassword(person.Password)

	return rs.Tx().Model(person).Insert()
}

func (dao *PersonDAO) UpdatePerson(rs app.RequestScope, id int, person *models.Person) error {
	oldPerson, err := dao.GetPerson(rs, id)
	if err != nil {
		return err
	}

	person.Id = id
	if person.Password == "" {
		person.Password = *&oldPerson.Password
	} else {
		person.Password = EncryptPassword(person.Password)
	}

	return rs.Tx().Model(person).Exclude("Id").Update()
}

func (dao *PersonDAO) DeletePerson(rs app.RequestScope, id int) error {
	person, err := dao.GetPerson(rs, id)
	if err != nil {
		return err
	}
	return rs.Tx().Model(person).Delete()
}

func EncryptPassword(password string) string {
	newPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(newPassword)
}
