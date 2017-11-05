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
	UpdateWallet(rs app.RequestScope, wallet_id int, wallet *models.Wallet) error
	//GetWalletByPersonId return the wallet from authenticated person
	GetWalletByPersonId(rs app.RequestScope, personId int) (models.Wallet, error)
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
func (s *WalletService) GetWallet(rs app.RequestScope, wallet_id int) (*models.Wallet, error) {
	wallet, err := s.dao.GetWallet(rs, wallet_id)
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

// GetWalletThrowVerification retorna a wallet que esse cartão faz parte.
// Deixa a verificação para ser feita onde esse metodo for invocado
func (s *WalletService) GetWalletThrowVerification(rs app.RequestScope, card_id int) (*models.Wallet, error) {
	wallet, err := s.dao.GetWallet(rs, card_id)
	if err != nil {
		return nil, err
	}

	return wallet, nil
}

func (s *WalletService) GetAuthenticatedPersonWallet(rs app.RequestScope) (models.Wallet, error) {
	return s.dao.GetWalletByPersonId(rs, rs.UserID())
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
func (s *WalletService) UpdateAuthenticatedPersonWallet(rs app.RequestScope, wallet *models.Wallet) (*models.Wallet, error) {
	wallet.PersonId = rs.UserID()

	oldWallet, err := s.GetAuthenticatedPersonWallet(rs)
	if err != nil {
		return nil, err
	}

	//card é usado pra representar a soma dos limites dos cartões
	card, err := NewCardService(daos.NewCardDAO()).GetWalletCardsLimits(rs, oldWallet.Id)
	if err != nil {
		return nil, err
	}

	wallet.Id = oldWallet.Id
	wallet.MaximumLimit = card.RealLimit
	wallet.AvaliableLimit = card.AvaliableLimit

	if wallet.CurrentLimit > wallet.AvaliableLimit {
		return nil, errors.NewAPIError(400, "ERROR", errors.Params{"message": "Attempted to increase to a higher limit than available. This is your avaliable limit: R$" + strconv.Itoa(int(card.AvaliableLimit)) + ". Pay some credit card to release more credit.", "developer_message": "current_limit greater than avaliable_limit of cards."})
	} else {
		wallet.AvaliableLimit = card.AvaliableLimit - wallet.CurrentLimit
	}

	if err := s.dao.UpdateWallet(rs, oldWallet.Id, wallet); err != nil {
		return nil, err
	}
	return s.dao.GetWallet(rs, oldWallet.Id)
}

func (s *WalletService) UpdateWalletLimits(rs app.RequestScope, card models.Card) error {
	wallet, err := s.GetWallet(rs, card.WalletId)
	if err != nil {
		return err
	}

	wallet.MaximumLimit = card.AvaliableLimit
	wallet.AvaliableLimit = card.AvaliableLimit - wallet.CurrentLimit

	return s.dao.UpdateWallet(rs, card.WalletId, wallet)
}
