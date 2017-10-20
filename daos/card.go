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
	err := rs.Tx().Select().Where(dbx.HashExp{"wallet_id": &wallet.Id}).OrderBy("cc_due_date DESC", "cc_current_limit ASC").All(&cards)
	return cards, err
}

// GetCardsByWalletId retorna uma lista de cartões de uma pessoa com id pespecífico.
func (dao *CardDAO) GetCardsByWalletId(rs app.RequestScope, personId int, walletId int) ([]models.Card, error) {
	cards := []models.Card{}
	person := models.Person{}
	errPerson := rs.Tx().Select().Where(dbx.HashExp{"id": personId}).One(&person)

	if errPerson != nil {
		return nil, errPerson
	}

	err := rs.Tx().Select().Where(dbx.HashExp{"person_id": &person.Id}).All(&cards)
	return cards, err
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

// Count retorna a quantidade de cartões.
func (dao *CardDAO) Count(rs app.RequestScope) (int, error) {
	var count int
	err := rs.Tx().Select("COUNT(*)").From("card").Row(&count)
	return count, err
}

//Query retorna os cartões no intervalo do offset e o limit.
func (dao *CardDAO) Query(rs app.RequestScope, offset, limit int) ([]models.Card, error) {
	cards := []models.Card{}
	err := rs.Tx().Select().OrderBy("id").Offset(int64(offset)).Limit(int64(limit)).All(&cards)
	return cards, err
}
