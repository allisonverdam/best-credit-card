package services

import (
	"strconv"

	"github.com/allisonverdam/best-credit-card/app"
	"github.com/allisonverdam/best-credit-card/daos"
	"github.com/allisonverdam/best-credit-card/errors"
	"github.com/allisonverdam/best-credit-card/models"
)

// WalletDAO specifies the interface of the wallet DAO needed by WalletService.
type WalletDAO interface {
	// GetWallet returns the wallet with the specified wallet ID.
	GetWallet(rs app.RequestScope, card_id int) (*models.Wallet, error)
	// CreateWallet saves a new wallet in the storage.
	CreateWallet(rs app.RequestScope, wallet *models.Wallet) error
	// UpdateWallet updates the wallet with given ID in the storage.
	UpdateWallet(rs app.RequestScope, card_id int, wallet *models.Wallet) error
	// DeleteWallet removes the wallet with given ID from the storage.
	DeleteWallet(rs app.RequestScope, card_id int) error
	//GetAuthenticatedPersonWallets return the wallets from authenticated person
	GetAuthenticatedPersonWallets(rs app.RequestScope, personId int) ([]models.Wallet, error)
}

// WalletService provides services related with wallets.
type WalletService struct {
	dao WalletDAO
}

// NewWalletService creates a new WalletService with the given wallet DAO.
func NewWalletService(dao WalletDAO) *WalletService {
	return &WalletService{dao}
}

// GetWallet returns the wallet with the specified the wallet ID.
func (s *WalletService) GetWallet(rs app.RequestScope, card_id int) (*models.Wallet, error) {
	wallet, err := s.dao.GetWallet(rs, card_id)
	if err != nil {
		return nil, err
	}

	//Verifica se o cartão pertence a pessoa que está autenticada
	err = VerifyPersonOwner(rs, wallet.PersonId, "wallet")
	if err != nil {
		return nil, err
	}
	return wallet, nil
}

func (s *WalletService) GetAuthenticatedPersonWallets(rs app.RequestScope) ([]models.Wallet, error) {
	return s.dao.GetAuthenticatedPersonWallets(rs, rs.UserID())
}

// CreateWallet creates a new wallet.
func (s *WalletService) CreateWallet(rs app.RequestScope, wallet *models.Wallet) (*models.Wallet, error) {
	if err := wallet.Validate(); err != nil {
		return nil, err
	}

	if rs.UserID() != 0 {
		//Verifica se o card_id da pessoa é igual o da que está autenticada
		err := VerifyPersonOwner(rs, wallet.PersonId, "person_id")
		if err != nil {
			return nil, err
		}
	}

	if err := s.dao.CreateWallet(rs, wallet); err != nil {
		return nil, err
	}
	return s.dao.GetWallet(rs, wallet.Id)
}

// UpdateWallet updates the wallet with the specified ID.
func (s *WalletService) UpdateWallet(rs app.RequestScope, card_id int, wallet *models.Wallet) (*models.Wallet, error) {
	if err := wallet.Validate(); err != nil {
		return nil, err
	}

	//Verifica se a carteira pertence a pessoa que está autenticada
	err := VerifyPersonOwner(rs, wallet.PersonId, "wallet")
	if err != nil {
		return nil, err
	}

	//card é usado pra representar a soma dos limites dos cartões
	card, err := NewCardService(daos.NewCardDAO()).GetWalletCardsLimits(rs, wallet.Id)
	if err != nil {
		return nil, err
	}

	wallet.MaximumLimit = card.RealLimit
	wallet.AvaliableLimit = card.AvaliableLimit

	if wallet.CurrentLimit > wallet.AvaliableLimit {
		return nil, errors.NewAPIError(400, "ERROR", errors.Params{"message": "Attempted to increase to a higher limit than available. This is your avaliable limit: R$" + strconv.Itoa(int(card.AvaliableLimit)) + ". Pay some credit card to release more credit.", "developer_message": "current_limit greater than avaliable_limit of cards."})
	} else {
		wallet.AvaliableLimit = card.AvaliableLimit - wallet.CurrentLimit
	}

	if err := s.dao.UpdateWallet(rs, card_id, wallet); err != nil {
		return nil, err
	}
	return s.dao.GetWallet(rs, card_id)
}

// DeleteWallet deletes the wallet with the specified ID.
func (s *WalletService) DeleteWallet(rs app.RequestScope, card_id int) (*models.Wallet, error) {
	wallet, err := s.dao.GetWallet(rs, card_id)
	if err != nil {
		return nil, err
	}

	//Verifica se a carteira pertence a pessoa que está autenticada
	err = VerifyPersonOwner(rs, wallet.PersonId, "wallet")
	if err != nil {
		return nil, err
	}

	err = s.dao.DeleteWallet(rs, card_id)
	return wallet, err
}

func (s *WalletService) UpdateWalletLimits(rs app.RequestScope, card models.Card) error {
	wallet, err := s.GetWallet(rs, card.WalletId)
	if err != nil {
		return err
	}

	//Verifica se a carteira pertence a pessoa que está autenticada
	err = VerifyPersonOwner(rs, wallet.PersonId, "wallet")
	if err != nil {
		return err
	}

	wallet.MaximumLimit = card.AvaliableLimit
	wallet.AvaliableLimit = card.AvaliableLimit - wallet.CurrentLimit

	return s.dao.UpdateWallet(rs, card.WalletId, wallet)
}
