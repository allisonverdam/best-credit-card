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
	rg.Get("/cards/<card_id>", r.get)
	rg.Get("/cards/wallets/<wallet_id>", r.cardsWallet)
	rg.Post("/cards", r.create)
	rg.Post("/cards/pay", r.payCreditCard)
	rg.Post("/cards/best-card", r.getBestCards)
	rg.Put("/cards/<id>", r.update)
	rg.Delete("/cards/<id>", r.delete)
}

func (r *cardResource) getBestCards(c *routing.Context) error {
	var order models.Order
	if err := c.Read(&order); err != nil {
		return err
	}
	if err := order.Validate(); err != nil {
		return err
	}

	card, err := r.service.GetBestCards(app.GetRequestScope(c), app.GetRequestScope(c).UserID(), &order)
	if err != nil {
		return err
	}

	return c.Write(card)
}

func (r *cardResource) payCreditCard(c *routing.Context) error {
	var order models.Order
	if err := c.Read(&order); err != nil {
		return err
	}
	if err := order.ValidateCardIdAndPrice(); err != nil {
		return err
	}

	card, err := r.service.PayCreditCard(app.GetRequestScope(c), *&order)
	if err != nil {
		return err
	}

	return c.Write(card)
}

func (r *cardResource) get(c *routing.Context) error {
	id, err := strconv.Atoi(c.Param("card_id"))
	if err != nil {
		return err
	}

	card, err := r.service.Get(app.GetRequestScope(c), id)
	if err != nil {
		return err
	}

	return c.Write(card)
}

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

func (r *cardResource) create(c *routing.Context) error {
	card := models.Card{}
	if err := c.Read(&card); err != nil {
		return err
	}

	if err := card.Validate(); err != nil {
		return err
	}

	cardBD, err := r.service.Create(app.GetRequestScope(c), &card)
	if err != nil {
		return err
	}

	return c.Write(cardBD)
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

	cardBD, err := r.service.Update(rs, id, card)
	if err != nil {
		return err
	}

	return c.Write(cardBD)
}

func (r *cardResource) delete(c *routing.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	card, err := r.service.Delete(app.GetRequestScope(c), id)
	if err != nil {
		return err
	}

	return c.Write(card)
}
