package services

import (
	"github.com/allisonverdam/best-credit-card/app"
	"github.com/allisonverdam/best-credit-card/daos"
	"github.com/allisonverdam/best-credit-card/models"
)

// personDAO specifies the interface of the person DAO needed by PersonService.
type personDAO interface {
	// Get returns the person with the specified person ID.
	Get(rs app.RequestScope, id int) (*models.Person, error)
	// GetWithoutPassword returns the person with the specified person ID whithout password.
	GetWithoutPassword(rs app.RequestScope, id int) (*models.Person, error)
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
	person, err := s.dao.GetWithoutPassword(rs, id)
	if err != nil {
		return nil, err
	}

	return person, nil
}

// Get returns the person with the specified the person ID.
func (s *PersonService) GetAuthenticatedPersonWallets(rs app.RequestScope, personId int) ([]models.Wallet, error) {
	walletDao := daos.NewWalletDAO()

	wallets, err := walletDao.GetAuthenticatedPersonWallets(rs, personId)
	if err != nil {
		return nil, err
	}

	return wallets, nil
}

// Update updates the person with the specified ID.
func (s *PersonService) Update(rs app.RequestScope, id int, person *models.Person) (*models.Person, error) {
	if err := s.dao.Update(rs, id, person); err != nil {
		return nil, err
	}
	return s.dao.GetWithoutPassword(rs, id)
}
