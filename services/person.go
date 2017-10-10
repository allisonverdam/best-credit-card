package services

import (
	"github.com/allisonverdam/best-credit-card/app"
	"github.com/allisonverdam/best-credit-card/models"
)

// personDAO specifies the interface of the person DAO needed by PersonService.
type personDAO interface {
	// Get returns the person with the specified person ID.
	Get(rs app.RequestScope, id int) (*models.Person, error)
	// Get returns the person with the specified person username.
	GetPersonByUserName(rs app.RequestScope, username string) (*models.Person, error)
	// Count returns the number of persons.
	Count(rs app.RequestScope) (int, error)
	// Query returns the list of persons with the given offset and limit.
	Query(rs app.RequestScope, offset, limit int) ([]models.Person, error)
	// Create saves a new person in the storage.
	Create(rs app.RequestScope, person *models.Person) error
	// Update updates the person with given ID in the storage.
	Update(rs app.RequestScope, id int, person *models.Person) error
	// Delete removes the person with given ID from the storage.
	Delete(rs app.RequestScope, id int) error
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

// Get returns the person with the specified the person username.
func (s *PersonService) GetPersonByUserName(rs app.RequestScope, username string) (*models.Person, error) {
	return s.dao.GetPersonByUserName(rs, username)
}

// Create creates a new person.
func (s *PersonService) Create(rs app.RequestScope, model *models.Person) (*models.Person, error) {
	if err := model.Validate(); err != nil {
		return nil, err
	}
	if err := s.dao.Create(rs, model); err != nil {
		return nil, err
	}
	return s.dao.Get(rs, model.Id)
}

// Update updates the person with the specified ID.
func (s *PersonService) Update(rs app.RequestScope, id int, model *models.Person) (*models.Person, error) {
	if err := model.Validate(); err != nil {
		return nil, err
	}
	if err := s.dao.Update(rs, id, model); err != nil {
		return nil, err
	}
	return s.dao.Get(rs, id)
}

// Delete deletes the person with the specified ID.
func (s *PersonService) Delete(rs app.RequestScope, id int) (*models.Person, error) {
	person, err := s.dao.Get(rs, id)
	if err != nil {
		return nil, err
	}
	err = s.dao.Delete(rs, id)
	return person, err
}

// Count returns the number of persons.
func (s *PersonService) Count(rs app.RequestScope) (int, error) {
	return s.dao.Count(rs)
}

// Query returns the persons with the specified offset and limit.
func (s *PersonService) Query(rs app.RequestScope, offset, limit int) ([]models.Person, error) {
	return s.dao.Query(rs, offset, limit)
}
