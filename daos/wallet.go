package daos

import (
	"github.com/allisonverdam/best-credit-card/app"
	"github.com/allisonverdam/best-credit-card/models"
	"github.com/go-ozzo/ozzo-dbx"
)

// WalletDAO faz a persistencia dos dados no bd
type WalletDAO struct{}

// NewWalletDAO cria um novo WalletDAO
func NewWalletDAO() *WalletDAO {
	return &WalletDAO{}
}

// Get reads the wallet with the specified ID from the database.
func (dao *WalletDAO) Get(rs app.RequestScope, id int) (*models.Wallet, error) {
	var wallet models.Wallet
	err := rs.Tx().Select().Model(id, &wallet)
	return &wallet, err
}

// Get reads the wallet with the specified ID from the database.
func (dao *WalletDAO) GetWalletByUserName(rs app.RequestScope, username string) (*models.Wallet, error) {
	var wallet models.Wallet
	err := rs.Tx().Select().Where(dbx.Like("username", username)).One(&wallet)
	return &wallet, err
}

// Create saves a new wallet record in the database.
// The Wallet.Id field will be populated with an automatically generated ID upon successful saving.
func (dao *WalletDAO) Create(rs app.RequestScope, wallet *models.Wallet) error {
	wallet.Id = 0
	return rs.Tx().Model(wallet).Insert()
}

// Update saves the changes to an wallet in the database.
func (dao *WalletDAO) Update(rs app.RequestScope, id int, wallet *models.Wallet) error {
	if _, err := dao.Get(rs, id); err != nil {
		return err
	}
	wallet.Id = id
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

// Count returns the number of the wallet records in the database.
func (dao *WalletDAO) Count(rs app.RequestScope) (int, error) {
	var count int
	err := rs.Tx().Select("COUNT(*)").From("wallet").Row(&count)
	return count, err
}

// Query retrieves the wallet records with the specified offset and limit from the database.
func (dao *WalletDAO) Query(rs app.RequestScope, offset, limit int) ([]models.Wallet, error) {
	wallets := []models.Wallet{}
	err := rs.Tx().Select().OrderBy("id").Offset(int64(offset)).Limit(int64(limit)).All(&wallets)
	return wallets, err
}
