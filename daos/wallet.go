package daos

import (
	"github.com/allisonverdam/best-credit-card/app"
	"github.com/allisonverdam/best-credit-card/models"
)

// WalletDAO faz a persistencia dos dados no bd
type WalletDAO struct{}

// NewWalletDAO cria um novo WalletDAO
func NewWalletDAO() *WalletDAO {
	return &WalletDAO{}
}

// Get reads the wallet with the specified ID from the database.
func (dao *WalletDAO) Get(rs app.RequestScope, id int) (*models.Wallet, error) {
	wallet := models.Wallet{}
	err := rs.Tx().Select().Model(id, &wallet)
	return &wallet, err
}

// Create saves a new wallet record in the database.
func (dao *WalletDAO) Create(rs app.RequestScope, wallet *models.Wallet) error {
	err := wallet.Validate()
	if err != nil {
		return err
	}
	return rs.Tx().Model(wallet).Insert()
}

// Update saves the changes to an wallet in the database.
func (dao *WalletDAO) Update(rs app.RequestScope, id int, wallet *models.Wallet) error {
	err := wallet.Validate()
	if err != nil {
		return err
	}

	if _, err := dao.Get(rs, id); err != nil {
		return err
	}
	return rs.Tx().Model(wallet).Exclude("Id").Update()
}

// Delete deletes an wallet with the specified ID from the database.
func (dao *WalletDAO) Delete(rs app.RequestScope, id int) error {
	wallet, err := dao.Get(rs, id)
	if err != nil {
		return err
	}
	return rs.Tx().Model(wallet).Delete()
}
