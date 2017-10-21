package services

import (
	"github.com/allisonverdam/best-credit-card/app"
	"github.com/allisonverdam/best-credit-card/models"
)

// personDAO specifies the interface of the person DAO needed by PersonService.
type personDAO interface {
	// Get returns the person with the specified person ID.
	Get(rs app.RequestScope, id int) (*models.Person, error)
	// Update updates the person with given ID in the storage.
	Update(rs app.RequestScope, id int, person *models.Person) error
	// Create saves a new person in the storage.
	Create(rs app.RequestScope, person *models.Person) error
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
func (s *PersonService) Get(rs app.RequestScope, id int) (*models.Person, error) {
	return s.dao.Get(rs, id)
}

// Update updates the person with the specified ID.
func (s *PersonService) Update(rs app.RequestScope, id int, person *models.Person) (*models.Person, error) {
	if err := person.ValidateOnUpdate(); err != nil {
		return nil, err
	}

	oldPerson, err := s.dao.Get(rs, id)
	if err != nil {
		return nil, err
	}

	person.Password = oldPerson.Password

	if err := s.dao.Update(rs, id, person); err != nil {
		return nil, err
	}
	return s.dao.Get(rs, id)
}

// Create creates a new person.
func (s *PersonService) Create(rs app.RequestScope, person *models.Person) (*models.Person, error) {
	if err := person.Validate(); err != nil {
		return nil, err
	}
	if err := s.dao.Create(rs, person); err != nil {
		return nil, err
	}
	return s.dao.Get(rs, person.Id)
}
