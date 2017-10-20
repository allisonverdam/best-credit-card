package services

import (
	"github.com/allisonverdam/best-credit-card/app"
	"github.com/allisonverdam/best-credit-card/models"
)

// WalletDAO specifies the interface of the wallet DAO needed by WalletService.
type WalletDAO interface {
	// Get returns the wallet with the specified wallet ID.
	Get(rs app.RequestScope, id int) (*models.Wallet, error)
	// Count returns the number of wallets.
	Count(rs app.RequestScope) (int, error)
	// Query returns the list of wallets with the given offset and limit.
	Query(rs app.RequestScope, offset, limit int) ([]models.Wallet, error)
	// Create saves a new wallet in the storage.
	Create(rs app.RequestScope, wallet *models.Wallet) error
	// Update updates the wallet with given ID in the storage.
	Update(rs app.RequestScope, id int, wallet *models.Wallet) error
	// Delete removes the wallet with given ID from the storage.
	Delete(rs app.RequestScope, id int) error
}

// WalletService provides services related with wallets.
type WalletService struct {
	dao WalletDAO
}

// NewWalletService creates a new WalletService with the given wallet DAO.
func NewWalletService(dao WalletDAO) *WalletService {
	return &WalletService{dao}
}

// Get returns the wallet with the specified the wallet ID.
func (s *WalletService) Get(rs app.RequestScope, id int) (*models.Wallet, error) {
	return s.dao.Get(rs, id)
}

// Create creates a new wallet.
func (s *WalletService) Create(rs app.RequestScope, model *models.Wallet) (*models.Wallet, error) {
	if err := model.Validate(); err != nil {
		return nil, err
	}
	if err := s.dao.Create(rs, model); err != nil {
		return nil, err
	}
	return s.dao.Get(rs, model.Id)
}

// Update updates the wallet with the specified ID.
func (s *WalletService) Update(rs app.RequestScope, id int, model *models.Wallet) (*models.Wallet, error) {
	if err := model.Validate(); err != nil {
		return nil, err
	}
	if err := s.dao.Update(rs, id, model); err != nil {
		return nil, err
	}
	return s.dao.Get(rs, id)
}

// Delete deletes the wallet with the specified ID.
func (s *WalletService) Delete(rs app.RequestScope, id int) (*models.Wallet, error) {
	wallet, err := s.dao.Get(rs, id)
	if err != nil {
		return nil, err
	}
	err = s.dao.Delete(rs, id)
	return wallet, err
}

// Count returns the number of wallets.
func (s *WalletService) Count(rs app.RequestScope) (int, error) {
	return s.dao.Count(rs)
}

// Query returns the wallets with the specified offset and limit.
func (s *WalletService) Query(rs app.RequestScope, offset, limit int) ([]models.Wallet, error) {
	return s.dao.Query(rs, offset, limit)
}
