package daos

import (
	"github.com/allisonverdam/best-credit-card/app"
	"github.com/allisonverdam/best-credit-card/models"
	dbx "github.com/go-ozzo/ozzo-dbx"
)

// WalletDAO faz a persistencia dos dados no bd
type WalletDAO struct{}

// NewWalletDAO cria um novo WalletDAO
func NewWalletDAO() *WalletDAO {
	return &WalletDAO{}
}

// GetWallet reads the wallet with the specified ID from the database.
func (dao *WalletDAO) GetWallet(rs app.RequestScope, id int) (*models.Wallet, error) {
	wallet := models.Wallet{}
	err := rs.Tx().Select().Model(id, &wallet)
	if err != nil {
		return nil, err
	}

	return &wallet, err
}

// GetAuthenticatedPersonWallets retorna todas as carteiras do usuario autenticado.
func (dao *WalletDAO) GetAuthenticatedPersonWallets(rs app.RequestScope, personId int) ([]models.Wallet, error) {
	wallet := []models.Wallet{}
	err := rs.Tx().Select().Where(dbx.HashExp{"person_id": personId}).All(&wallet)
	return wallet, err
}

// CreateWallet saves a new wallet record in the database.
func (dao *WalletDAO) CreateWallet(rs app.RequestScope, wallet *models.Wallet) error {
	wallet.CurrentLimit = 0
	wallet.MaximumLimit = 0

	return rs.Tx().Model(wallet).Insert()
}

// UpdateWallet saves the changes to an wallet in the database.
func (dao *WalletDAO) UpdateWallet(rs app.RequestScope, id int, wallet *models.Wallet) error {
	if _, err := dao.GetWallet(rs, id); err != nil {
		return err
	}
	return rs.Tx().Model(wallet).Exclude("Id").Update()
}

// DeleteWallet deletes an wallet with the specified ID from the database.
func (dao *WalletDAO) DeleteWallet(rs app.RequestScope, id int) error {
	wallet, err := dao.GetWallet(rs, id)
	if err != nil {
		return err
	}
	return rs.Tx().Model(wallet).Delete()
}
