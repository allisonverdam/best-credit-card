package services

import (
	"net/http"
	"sort"
	"time"

	"github.com/allisonverdam/best-credit-card/daos"

	"github.com/allisonverdam/best-credit-card/app"
	"github.com/allisonverdam/best-credit-card/errors"
	"github.com/allisonverdam/best-credit-card/models"
)

// cardDAO specifies the interface of the card DAO needed by CardService.
type cardDAO interface {
	// GetCard returns the card with the specified ID.
	GetCard(rs app.RequestScope, card_id int) (*models.Card, error)
	// GetBestCardsByWallet returns the best cards to a order in a wallet.
	GetBestCardsByWallet(rs app.RequestScope, wallet models.Wallet) ([]models.Card, error)
	// GetCardsByWallet returns the cards of a wallet.
	GetCardsByWallet(rs app.RequestScope, wallet models.Wallet) ([]models.Card, error)
	//GetWalletCardsLimits return the limits of a wallet with the specified ID
	GetWalletCardsLimits(rs app.RequestScope, walletId int) (*models.Card, error)
	// CreateCard saves a new card in the storage.
	CreateCard(rs app.RequestScope, card *models.Card) error
	// UpdateCard updates the card with given ID in the storage.
	UpdateCard(rs app.RequestScope, card_id int, card *models.Card) error
	// DeleteCard removes the card with given ID from the storage.
	DeleteCard(rs app.RequestScope, card_id int) error
}

// CardService provides services related with cards.
type CardService struct {
	dao cardDAO
}

type By func(c1, c2 *models.Card) bool

type cardSorter struct {
	cards []models.Card
	by    func(p1, p2 *models.Card) bool
}

// NewCardService creates a new CardService with the given card DAO.
func NewCardService(dao cardDAO) *CardService {
	return &CardService{dao}
}

// GetCard returns the card with the specified the card ID.
func (s *CardService) GetCard(rs app.RequestScope, card_id int) (*models.Card, error) {
	card, err := s.dao.GetCard(rs, card_id)
	if err != nil {
		return nil, err
	}

	_, walletErr := NewWalletService(daos.NewWalletDAO()).GetWallet(rs, card.WalletId)
	if walletErr != nil {
		return nil, walletErr
	}

	return card, err
}

// GetCard returns the card with the specified the card ID.
func (s *CardService) PayCreditCard(rs app.RequestScope, order models.Order) (*models.Card, error) {
	card, err := s.dao.GetCard(rs, order.CardId)
	if err != nil {
		return nil, err
	}

	_, walletErr := NewWalletService(daos.NewWalletDAO()).GetWallet(rs, card.WalletId)
	if walletErr != nil {
		return nil, walletErr
	}

	if (card.AvaliableLimit + order.Price) > card.RealLimit {
		return nil, errors.NewAPIError(http.StatusBadRequest, "ERROR", errors.Params{"message": "Price greater than the maximum card limit.", "developer_message": ""})

	}

	card.AvaliableLimit += order.Price

	if err := s.dao.UpdateCard(rs, card.Id, card); err != nil {
		return nil, err
	}

	if err := s.UpdateWalletLimits(rs, card.WalletId); err != nil {
		return nil, err
	}

	return card, nil
}

// GetBestCards returns the best cards for a order in a wallet.
func (s *CardService) GetBestCards(rs app.RequestScope, order *models.Order) ([]models.Card, error) {

	wallet, err := NewWalletService(daos.NewWalletDAO()).GetAuthenticatedPersonWallet(rs)
	if err != nil {
		return nil, err
	}

	cards, err := s.dao.GetBestCardsByWallet(rs, *&wallet)
	if err != nil {
		return nil, err
	}

	bestCards := []models.Card{}

	price := *&order.Price
	today := time.Now().Day()

	//verificando se o dueDate é menor que o dia atual
	//assim ele é mais longe, então adiciono 100 ao valor para fica maior
	for i, card := range cards {
		if card.DueDate <= today {
			cards[i].DueDate += 100
		}
	}

	// definindo como sera ordenado.
	dueDate := func(c1, c2 *models.Card) bool {
		return c1.DueDate > c2.DueDate
	}

	//ordenando pelo maior dueDate
	By(dueDate).Sort(cards)

	for _, card := range cards {
		if price > card.AvaliableLimit {
			bestCards = append(bestCards, card)
			price -= card.AvaliableLimit
		} else {
			bestCards = append(bestCards, card)
			price -= card.AvaliableLimit
			break
		}
	}

	if price > 0 {
		return nil, errors.NewAPIError(http.StatusBadRequest, "ERROR", errors.Params{"message": "You don't have enough credit to make this purchase.", "developer_message": ""})
	}

	return bestCards, nil
}

// GetCard returns the card with the specified the card ID.
func (s *CardService) GetAuthenticatedPersonCards(rs app.RequestScope) ([]models.Card, error) {
	wallet, err := NewWalletService(daos.NewWalletDAO()).GetAuthenticatedPersonWallet(rs)
	if err != nil {
		return nil, err
	}

	return s.dao.GetCardsByWallet(rs, wallet)
}

// CreateCard creates a new card.
func (s *CardService) CreateCard(rs app.RequestScope, card *models.Card) (*models.Card, error) {
	if validationErr := card.Validate(); validationErr != nil {
		return nil, validationErr
	}

	wallet, err := NewWalletService(daos.NewWalletDAO()).GetAuthenticatedPersonWallet(rs)
	if err != nil {
		return nil, err
	}

	card.WalletId = wallet.Id
	if err := s.dao.CreateCard(rs, card); err != nil {
		return nil, err
	}

	if err := s.UpdateWalletLimits(rs, wallet.Id); err != nil {
		return nil, err
	}

	return s.dao.GetCard(rs, card.Id)
}

// UpdateCard updates the card with the specified ID.
func (s *CardService) UpdateCard(rs app.RequestScope, card_id int, card *models.Card) (*models.Card, error) {
	old_card, err := s.GetCard(rs, card_id)
	if err != nil {
		return nil, err
	}

	if err := s.dao.UpdateCard(rs, old_card.Id, card); err != nil {
		return nil, err
	}

	if err := s.UpdateWalletLimits(rs, card.WalletId); err != nil {
		return nil, err
	}

	return s.dao.GetCard(rs, card_id)
}

// DeleteCard deletes the card with the specified ID.
func (s *CardService) DeleteCard(rs app.RequestScope, card_id int) (*models.Card, error) {
	card, err := s.GetCard(rs, card_id)
	if err != nil {
		return nil, err
	}

	if cardErr := s.dao.DeleteCard(rs, card_id); cardErr != nil {
		return nil, cardErr
	}

	if err := s.UpdateWalletLimits(rs, card.WalletId); err != nil {
		return nil, err
	}

	return card, nil
}

//Retorna os limites dos cartoes de uma carteira
func (s *CardService) GetWalletCardsLimits(rs app.RequestScope, walletId int) (*models.Card, error) {
	_, err := NewWalletService(daos.NewWalletDAO()).GetWallet(rs, walletId)
	if err != nil {
		return nil, err
	}

	card, err := s.dao.GetWalletCardsLimits(rs, walletId)
	if err != nil {
		return nil, err
	}

	return card, nil
}

func (s *CardService) UpdateWalletLimits(rs app.RequestScope, walletId int) error {
	card, err := s.GetWalletCardsLimits(rs, walletId)
	if err != nil {
		return err
	}

	return NewWalletService(daos.NewWalletDAO()).UpdateWalletLimits(rs, *card)
}

func (by By) Sort(cards []models.Card) {
	ps := &cardSorter{
		cards: cards,
		by:    by, // Metodo definido para fazer a ordenação
	}
	sort.Sort(ps)
}

// Len is part of sort.Interface.
func (s *cardSorter) Len() int {
	return len(s.cards)
}

// Swap is part of sort.Interface.
func (s *cardSorter) Swap(i, j int) {
	s.cards[i], s.cards[j] = s.cards[j], s.cards[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s *cardSorter) Less(i, j int) bool {
	return s.by(&s.cards[i], &s.cards[j])
}
