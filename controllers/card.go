// @SubApi Card management API [/cards]
package controllers

import (
	"strconv"

	"github.com/allisonverdam/best-credit-card/app"
	"github.com/allisonverdam/best-credit-card/models"
	routing "github.com/go-ozzo/ozzo-routing"
)

type (
	// cardService especifica a interface que é utilizada pelo cardResource
	cardService interface {
		Get(rs app.RequestScope, id int) (*models.Card, error)
		GetBestCards(rs app.RequestScope, personId int, order *models.Order) ([]models.Card, error)
		GetCardsByWalletId(rs app.RequestScope, personId int, walletId int) ([]models.Card, error)
		PayCreditCard(rs app.RequestScope, order models.Order) (*models.Card, error)
		Create(rs app.RequestScope, card *models.Card) (*models.Card, error)
		Update(rs app.RequestScope, id int, card *models.Card) (*models.Card, error)
		Delete(rs app.RequestScope, id int) (*models.Card, error)
	}

	// cardResource define os handlers para as chamadas do controller.
	cardResource struct {
		service cardService
	}
)

// ServeCard define as rotas.
func ServeCardResource(rg *routing.RouteGroup, service cardService) {
	r := &cardResource{service}
	rg.Get("/cards/<id>", r.get)
	rg.Get("/cards/wallets/<wallet_id>", r.cardsWallet)
	rg.Post("/cards", r.create)
	rg.Post("/cards/pay", r.payCreditCard)
	rg.Post("/cards/best-card", r.getBestCards)
	rg.Put("/cards/<id>", r.update)
	rg.Delete("/cards/<id>", r.delete)
}

// @Title getBestCards
// @Description Retorna o melhor cartão para a compra.
// @Accept  json
// @Param   price     body    int     true        "Valor da compra."
// @Param   wallet_id        body   int     true        "ID da 'Wallet' onde vai buscar o(s) melhor(es) cartões."
// @Success 200 {array}  models.Card "Retorna uma lista contendo o(s) melhor(es) cartões para essa compra."
// @Failure 403 {object} errors.APIError    "O parametro 'wallet_id' informado não pertence ao usuário autenticado."
// @Failure 400 {object} errors.APIError    "O parametro 'price' deve ser maior que 0."
// @Resource /cards
// @Router /cards/best-card/ [get]
func (r *cardResource) getBestCards(c *routing.Context) error {
	var order models.Order
	if err := c.Read(&order); err != nil {
		return err
	}
	if err := order.Validate(); err != nil {
		return err
	}

	response, err := r.service.GetBestCards(app.GetRequestScope(c), app.GetRequestScope(c).UserID(), &order)
	if err != nil {
		return err
	}

	return c.Write(response)
}

// @Title payCreditCard
// @Description Paga um cartao para liberar crédito.
// @Accept  json
// @Param   price     body    int     true        "Valor da compra."
// @Param   card_id        body   int     true        "ID do 'Card' no qual vai efetuar o pagamento."
// @Success 200 {object}  models.Card "Retorna os dados do cartão após efetuar o pagamento."
// @Failure 403 {object} errors.APIError    "O parametro 'card_id' informado não pertence ao usuário autenticado."
// @Failure 400 {object} errors.APIError    "O parametro 'price' deve ser maior que 0."
// @Resource /cards
// @Router /cards/pay/ [post]
func (r *cardResource) payCreditCard(c *routing.Context) error {
	var order models.Order
	if err := c.Read(&order); err != nil {
		return err
	}
	if err := order.ValidateCardIdAndPrice(); err != nil {
		return err
	}

	response, err := r.service.PayCreditCard(app.GetRequestScope(c), *&order)
	if err != nil {
		return err
	}

	return c.Write(response)
}

func (r *cardResource) get(c *routing.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	response, err := r.service.Get(app.GetRequestScope(c), id)
	if err != nil {
		return err
	}

	return c.Write(response)
}

// @Title cardsWallet
// @Description Retorna a lista de 'Cards' de uma determinada 'Wallet'.
// @Accept  json
// @Param   wallet_id     path    int     true        "ID da 'Wallet' que desejamos buscar os 'Cards'."
// @Success 200 {array}  models.Card "Retorna os 'Cards' que fazem parte dessa 'Wallet'."
// @Failure 403 {object} errors.APIError    "O parametro 'card_id' informado não pertence ao usuário autenticado."
// @Failure 400 {object} errors.APIError    "O parametro 'price' deve ser maior que 0."
// @Resource /cards
// @Router /cards/wallets/{wallet_id} [get]
func (r *cardResource) cardsWallet(c *routing.Context) error {
	wallet_id, err := strconv.Atoi(c.Param("wallet_id"))
	if err != nil {
		return err
	}

	rs := app.GetRequestScope(c)

	cards, err := r.service.GetCardsByWalletId(rs, rs.UserID(), wallet_id)
	if err != nil {
		return err
	}

	return c.Write(cards)
}

// @Title create
// @Description Cria um novo 'Card'.
// @Accept  json
// @Param   number     body    float64     true        "Número do cartão."
// @Param   due_date     body    int     true        "Data de vencimento do cartão."
// @Param   expiration_month     body    int     true	"teste"
// @Param   expiration_year     body    int	true	"teste"
// @Param   cvv     body    int     true	"teste"
// @Param   real_limit     body    float64     true	"teste"
// @Param   current_limit     body    float64     true	"teste"
// @Param   wallet_id     body    int     true	"teste"
// @Success 200 {object}  models.Card "Retorna o cartão que acabou de ser criado."
// @Failure 403 {object} errors.APIError    "O parametro 'wallet_id' informado não pertence ao usuário autenticado."
// @Failure 400 {object} errors.APIError    "Se estiver faltando algum parametro."
// @Failure 400 {object} errors.APIError    "O parametro 'current_limit' deve ser menor ou igual ao 'real_limit'."
// @Resource /cards
// @Router /cards/ [post]
func (r *cardResource) create(c *routing.Context) error {
	var card models.Card
	if err := c.Read(&card); err != nil {
		return err
	}

	if err := card.Validate(); err != nil {
		return err
	}

	response, err := r.service.Create(app.GetRequestScope(c), &card)
	if err != nil {
		return err
	}

	return c.Write(response)
}

func (r *cardResource) update(c *routing.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	rs := app.GetRequestScope(c)

	card, err := r.service.Get(rs, id)
	if err != nil {
		return err
	}

	if err := c.Read(card); err != nil {
		return err
	}

	response, err := r.service.Update(rs, id, card)
	if err != nil {
		return err
	}

	return c.Write(response)
}

func (r *cardResource) delete(c *routing.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	response, err := r.service.Delete(app.GetRequestScope(c), id)
	if err != nil {
		return err
	}

	return c.Write(response)
}
