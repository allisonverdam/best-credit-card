package daos

import (
	"github.com/allisonverdam/best-credit-card/app"
	"github.com/allisonverdam/best-credit-card/models"
	dbx "github.com/go-ozzo/ozzo-dbx"
)

// CardDAO faz a persistencia dos dados no bd.
type CardDAO struct{}

// NewCardDAO cria um novo CardDAO.
func NewCardDAO() *CardDAO {
	return &CardDAO{}
}

// GetCard retorna um cartão com id específico.
func (dao *CardDAO) GetCard(rs app.RequestScope, id int) (*models.Card, error) {
	card := models.Card{}
	err := rs.Tx().Select().Model(id, &card)
	return &card, err
}

// GetCardsByWalletId retorna uma lista de cartões de uma pessoa com id pespecífico.
func (dao *CardDAO) GetBestCardsByWallet(rs app.RequestScope, personId int, wallet models.Wallet) ([]models.Card, error) {
	cards := []models.Card{}

	//pega os cartões de uma determinada carteira, e ordena pelo maior cc_due_date
	//caso tenha cartões com o cc_due_date igual retorna o com menor limite primeiro
	if err := rs.Tx().Select().Where(dbx.HashExp{"wallet_id": &wallet.Id}).OrderBy("cc_due_date DESC", "cc_avaliable_limit ASC").All(&cards); err != nil {
		return nil, err
	}

	return cards, nil
}

// GetCardsByWalletId retorna uma lista de cartões de uma pessoa com id pespecífico.
func (dao *CardDAO) GetCardsByWallet(rs app.RequestScope, personId int, wallet models.Wallet) ([]models.Card, error) {
	cards := []models.Card{}

	errQuery := rs.Tx().Select().Where(dbx.HashExp{"wallet_id": &wallet.Id}).All(&cards)
	if errQuery != nil {
		return nil, errQuery
	}

	return cards, nil
}

// Create salva um novo cartão.
func (dao *CardDAO) CreateCard(rs app.RequestScope, card *models.Card) error {
	return rs.Tx().Model(card).Insert()
}

// UpdateCard atualiza os dados de um catrão com id específico.
func (dao *CardDAO) UpdateCard(rs app.RequestScope, id int, card *models.Card) error {
	if _, err := dao.GetCard(rs, id); err != nil {
		return err
	}
	card.Id = id
	return rs.Tx().Model(card).Exclude("Id").Update()
}

// Query retrieves the card records with the specified offset and limit from the database.
// DeleteCard deleta um cartão com id específico.
func (dao *CardDAO) DeleteCard(rs app.RequestScope, id int) error {
	card, err := dao.GetCard(rs, id)
	if err != nil {
		return err
	}
	return rs.Tx().Model(card).Delete()
}

// GetWalletCardsLimits retorna a soma do limite real e do limite atual de todos os cartoes de uma carteira.
func (dao *CardDAO) GetWalletCardsLimits(rs app.RequestScope, walletId int) (*models.Card, error) {
	card := models.Card{}

	errQuery := rs.Tx().Select("COALESCE(SUM(cc_real_limit), 0) cc_real_limit, COALESCE(SUM(cc_avaliable_limit), 0) cc_avaliable_limit").Where(dbx.HashExp{"wallet_id": walletId}).One(&card)
	if errQuery != nil {
		return nil, errQuery
	}

	//Adicionando wallet_id ao objeto, porque não vinha na query
	card.WalletId = walletId
	return &card, nil
}
