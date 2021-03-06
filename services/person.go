package services

import (
	"github.com/allisonverdam/best-credit-card/app"
	"github.com/allisonverdam/best-credit-card/models"
)

// personDAO specifies the interface of the person DAO needed by PersonService.
type personDAO interface {
	// Get returns the person with the specified person ID.
	GetPerson(rs app.RequestScope, id int) (*models.Person, error)
	// GetPersonWithoutPassword returns the person with the specified person ID whithout password.
	GetPersonWithoutPassword(rs app.RequestScope, id int) (*models.Person, error)
	// UpdatePerson updates the person with given ID in the storage.
	UpdatePerson(rs app.RequestScope, id int, person *models.Person) error
	// Create saves a new person in the storage.
	CreatePerson(rs app.RequestScope, person *models.Person) error
}

// PersonService provides services related with persons.
type PersonService struct {
	dao personDAO
}

// NewPersonService creates a new PersonService with the given person DAO.
func NewPersonService(dao personDAO) *PersonService {
	return &PersonService{dao}
}

// Get returns the person with the specified the person ID.
func (s *PersonService) GetPerson(rs app.RequestScope, id int) (*models.Person, error) {
	if err := VerifyPersonOwner(rs, id, "person_id"); err != nil {
		return nil, err
	}

	person, err := s.dao.GetPersonWithoutPassword(rs, id)
	if err != nil {
		return nil, err
	}

	return person, nil
}

// UpdateAuthenticatedPerson updates the person with the specified ID.
func (s *PersonService) UpdateAuthenticatedPerson(rs app.RequestScope, person *models.Person) (*models.Person, error) {
	id := rs.UserID()
	if err := s.dao.UpdatePerson(rs, id, person); err != nil {
		return nil, err
	}
	return s.dao.GetPersonWithoutPassword(rs, id)
}
