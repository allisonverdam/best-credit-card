package services

import (
	"net/http"

	"github.com/allisonverdam/best-credit-card/daos"

	"github.com/allisonverdam/best-credit-card/app"
	"github.com/allisonverdam/best-credit-card/errors"
	"github.com/allisonverdam/best-credit-card/models"
)

// cardDAO specifies the interface of the card DAO needed by CardService.
type cardDAO interface {
	// Get returns the card with the specified card ID.
	Get(rs app.RequestScope, id int) (*models.Card, error)
	// Get returns the card with the specified card ID.
	GetBestCardsByWalletId(rs app.RequestScope, personId int, walletId int) ([]models.Card, error)
	// Get returns the card with the specified card ID.
	GetCardsByWalletId(rs app.RequestScope, personId int, walletId int) ([]models.Card, error)
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
	card, err := s.dao.Get(rs, id)

	wallet, walletErr := GetWalletByCard(rs, card)
	if walletErr != nil {
		return nil, walletErr
	}

	//Verifica se o cartão pertence a pessoa que está autenticada
	err = VerifyPersonOwner(rs, wallet.PersonId, "card")
	if err != nil {
		return nil, err
	}

	return card, err
}

// Get returns the card with the specified the card ID.
func (s *CardService) PayCreditCard(rs app.RequestScope, order models.Order) (*models.Card, error) {
	if err := order.ValidateCardIdAndPrice(); err != nil {
		return nil, err
	}

	card, err := s.dao.Get(rs, order.CardId)
	if err != nil {
		return nil, err
	}

	wallet, err := GetWalletByCard(rs, card)
	if err != nil {
		return nil, err
	}

	person, err := GetPersonByWallet(rs, wallet)
	if err != nil {
		return nil, err
	}

	//Verifica se o cartão pertence a pessoa que está autenticada
	err = VerifyPersonOwner(rs, person.Id, "card")
	if err != nil {
		return nil, err
	}

	if (card.CurrentLimit + order.Price) > card.RealLimit {
		return nil, errors.NewAPIError(http.StatusPreconditionFailed, "ERROR", errors.Params{"message": "Price greater than the maximum card limit."})

	}

	card.CurrentLimit += order.Price

	if err := s.dao.Update(rs, card.Id, card); err != nil {
		return nil, err
	}

	return card, nil
}

// Get returns the card with the specified the card ID.
func (s *CardService) GetBestCards(rs app.RequestScope, personId int, order *models.Order) ([]models.Card, error) {
	if err := order.Validate(); err != nil {
		return nil, err
	}

	cards, err := s.dao.GetBestCardsByWalletId(rs, personId, order.WalletId)
	if err != nil {
		return nil, err
	}

	bestCards := []models.Card{}

	price := *&order.Price

	for _, card := range cards {
		if price <= 0 {
			break
		}

		if price > card.RealLimit {
			bestCards = append(bestCards, card)
			price -= card.RealLimit
		} else {
			bestCards = append(bestCards, card)
			break
		}
	}

	return bestCards, err
}

// Get returns the card with the specified the card ID.
func (s *CardService) GetCardsByWalletId(rs app.RequestScope, personId int, walletId int) ([]models.Card, error) {
	return s.dao.GetCardsByWalletId(rs, personId, walletId)
}

// Create creates a new card.
func (s *CardService) Create(rs app.RequestScope, card *models.Card) (*models.Card, error) {
	wallet, err := GetWalletByCard(rs, card)
	if err != nil {
		return nil, err
	}

	//Verifica se a carteira pertence a pessoa que está autenticada
	VerifyPersonOwner(rs, wallet.PersonId, "wallet")

	if err := s.dao.Create(rs, card); err != nil {
		return nil, err
	}
	return s.dao.Get(rs, card.Id)
}

// Update updates the card with the specified ID.
func (s *CardService) Update(rs app.RequestScope, id int, card *models.Card) (*models.Card, error) {
	if err := s.dao.Update(rs, id, card); err != nil {
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

//Retorna a carteira que o cartão faz parte
func GetWalletByCard(rs app.RequestScope, card *models.Card) (*models.Wallet, error) {
	walletDao := daos.NewWalletDAO()
	wallet, err := NewWalletService(walletDao).Get(rs, card.WalletId)

	return wallet, err
}

//Retorna a pessoa dona do cartão
func GetPersonByWallet(rs app.RequestScope, wallet *models.Wallet) (*models.Person, error) {
	personDao := daos.NewPersonDAO()
	person, err := NewPersonService(personDao).Get(rs, wallet.Id)

	return person, err
}

func VerifyPersonOwner(rs app.RequestScope, id int, resource string) error {
	if rs.UserID() != id {
		return errors.NewAPIError(http.StatusForbidden, "FORBIDDEN", errors.Params{"message": "This " + resource + " does not belong to this user."})
	}
	return nil
}
