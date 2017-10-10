package services

import (
	"github.com/allisonverdam/best-credit-card/app"
	"github.com/allisonverdam/best-credit-card/models"
)

// cardDAO specifies the interface of the card DAO needed by CardService.
type cardDAO interface {
	// Get returns the card with the specified card ID.
	Get(rs app.RequestScope, id int) (*models.Card, error)
	// Get returns the card with the specified card ID.
	GetCardsByPersonId(rs app.RequestScope, personId int) ([]models.Card, error)
	// Count returns the number of cards.
	Count(rs app.RequestScope) (int, error)
	// Query returns the list of cards with the given offset and limit.
	Query(rs app.RequestScope, offset, limit int) ([]models.Card, error)
	// Create saves a new card in the storage.
	Create(rs app.RequestScope, card *models.Card) error
	// Update updates the card with given ID in the storage.
	Update(rs app.RequestScope, id int, card *models.Card) error
	// Delete removes the card with given ID from the storage.
	Delete(rs app.RequestScope, id int) error
}

// CardService provides services related with cards.
type CardService struct {
	dao cardDAO
}

// NewCardService creates a new CardService with the given card DAO.
func NewCardService(dao cardDAO) *CardService {
	return &CardService{dao}
}

// Get returns the card with the specified the card ID.
func (s *CardService) Get(rs app.RequestScope, id int) (*models.Card, error) {
	return s.dao.Get(rs, id)
}

// Get returns the card with the specified the card ID.
func (s *CardService) GetCardsByPersonId(rs app.RequestScope, personId int) ([]models.Card, error) {
	return s.dao.GetCardsByPersonId(rs, personId)
}

// Create creates a new card.
func (s *CardService) Create(rs app.RequestScope, model *models.Card) (*models.Card, error) {
	if err := model.Validate(); err != nil {
		return nil, err
	}
	if err := s.dao.Create(rs, model); err != nil {
		return nil, err
	}
	return s.dao.Get(rs, model.Id)
}

// Update updates the card with the specified ID.
func (s *CardService) Update(rs app.RequestScope, id int, model *models.Card) (*models.Card, error) {
	if err := model.Validate(); err != nil {
		return nil, err
	}
	if err := s.dao.Update(rs, id, model); err != nil {
		return nil, err
	}
	return s.dao.Get(rs, id)
}

// Delete deletes the card with the specified ID.
func (s *CardService) Delete(rs app.RequestScope, id int) (*models.Card, error) {
	card, err := s.dao.Get(rs, id)
	if err != nil {
		return nil, err
	}
	err = s.dao.Delete(rs, id)
	return card, err
}

// Count returns the number of cards.
func (s *CardService) Count(rs app.RequestScope) (int, error) {
	return s.dao.Count(rs)
}

// Query returns the cards with the specified offset and limit.
func (s *CardService) Query(rs app.RequestScope, offset, limit int) ([]models.Card, error) {
	return s.dao.Query(rs, offset, limit)
}
