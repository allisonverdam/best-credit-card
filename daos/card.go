package daos

import (
	"net/http"

	"github.com/allisonverdam/best-credit-card/app"
	"github.com/allisonverdam/best-credit-card/errors"
	"github.com/allisonverdam/best-credit-card/models"
	dbx "github.com/go-ozzo/ozzo-dbx"
)

// CardDAO faz a persistencia dos dados no bd.
type CardDAO struct{}

// NewCardDAO cria um novo CardDAO.
func NewCardDAO() *CardDAO {
	return &CardDAO{}
}

// Get retorna um cartão com id específico.
func (dao *CardDAO) Get(rs app.RequestScope, id int) (*models.Card, error) {
	card := models.Card{}
	err := rs.Tx().Select().Model(id, &card)
	return &card, err
}

// GetCardsByWalletId retorna uma lista de cartões de uma pessoa com id pespecífico.
func (dao *CardDAO) GetBestCardsByWalletId(rs app.RequestScope, personId int, walletId int) ([]models.Card, error) {
	cards := []models.Card{}
	wallet := models.Wallet{}

	//verifica se a carteira existe
	errWallet := rs.Tx().Select().Where(dbx.HashExp{"id": walletId}).One(&wallet)
	if errWallet != nil {
		return nil, errWallet
	}

	//Verifica se a carteira pertence a pessoa que está autenticada
	if *&wallet.PersonId != personId {
		return nil, errors.NewAPIError(http.StatusForbidden, "FORBIDDEN", errors.Params{"message": "This wallet does not belong to this user."})
	}

	//pega os cartões de uma determinada carteira, e ordena pelo maior cc_due_date
	//caso tenha cartões com o cc_due_date igual retorna o com menor limite primeiro
	rs.Tx().Select().Where(dbx.HashExp{"wallet_id": &wallet.Id}).OrderBy("cc_due_date DESC", "cc_current_limit ASC").All(&cards)

	return cards, nil
}

// GetCardsByWalletId retorna uma lista de cartões de uma pessoa com id pespecífico.
func (dao *CardDAO) GetCardsByWalletId(rs app.RequestScope, personId int, walletId int) ([]models.Card, error) {
	cards := []models.Card{}
	wallet := models.Wallet{}

	errWallet := rs.Tx().Select().Where(dbx.HashExp{"id": walletId}).One(&wallet)
	if errWallet != nil {
		return nil, errWallet
	}

	if *&wallet.PersonId != personId {
		return nil, errors.NewAPIError(http.StatusForbidden, "FORBIDDEN", errors.Params{"message": "This wallet does not belong to the authenticated user."})
	}

	errQuery := rs.Tx().Select().Where(dbx.HashExp{"wallet_id": &wallet.Id}).All(&cards)
	if errQuery != nil {
		return nil, errQuery
	}

	return cards, nil
}

// Create salva um novo cartão.
func (dao *CardDAO) Create(rs app.RequestScope, card *models.Card) error {
	return rs.Tx().Model(card).Insert()
}

// Update atualiza os dados de um catrão com id específico.
func (dao *CardDAO) Update(rs app.RequestScope, id int, card *models.Card) error {
	if _, err := dao.Get(rs, id); err != nil {
		return err
	}
	card.Id = id
	return rs.Tx().Model(card).Exclude("Id").Update()
}

// Query retrieves the card records with the specified offset and limit from the database.
// Delete deleta um cartão com id específico.
func (dao *CardDAO) Delete(rs app.RequestScope, id int) error {
	card, err := dao.Get(rs, id)
	if err != nil {
		return err
	}
	return rs.Tx().Model(card).Delete()
}

// GetWalletCardsLimits retorna a soma do limite real e do limite atual de todos os cartoes de uma carteira.
func (dao *CardDAO) GetWalletCardsLimits(rs app.RequestScope, id int) (*models.Card, error) {
	cardLimit := models.Card{}
	card, err := dao.Get(rs, id)
	if err != nil {
		return nil, err
	}
	errQuery := rs.Tx().Select("SUM(cc_real_limit) cc_real_limit, SUM(cc_current_limit) cc_current_limit").Where(dbx.HashExp{"wallet_id": &card.WalletId}).All(&cardLimit)
	if errQuery != nil {
		return nil, errQuery
	}
	return &cardLimit, nil
}
