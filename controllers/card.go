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
		GetCardsByPersonId(rs app.RequestScope, personId int) ([]models.Card, error)
		Query(rs app.RequestScope, offset, limit int) ([]models.Card, error)
		Count(rs app.RequestScope) (int, error)
		Create(rs app.RequestScope, model *models.Card) (*models.Card, error)
		Update(rs app.RequestScope, id int, model *models.Card) (*models.Card, error)
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
	rg.Get("/cards", r.query)
	rg.Post("/cards", r.create)
	rg.Post("/bestCard", r.getBestCard)
	rg.Put("/cards/<id>", r.update)
	rg.Delete("/cards/<id>", r.delete)
}

// Função que retorna o melhor cartão para a compra
func (r *cardResource) getBestCard(c *routing.Context) error {
	response, err := r.service.GetCardsByPersonId(app.GetRequestScope(c), app.GetRequestScope(c).UserID())
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

func (r *cardResource) query(c *routing.Context) error {
	rs := app.GetRequestScope(c)
	count, err := r.service.Count(rs)
	if err != nil {
		return err
	}
	paginatedList := getPaginatedListFromRequest(c, count)
	items, err := r.service.Query(app.GetRequestScope(c), paginatedList.Offset(), paginatedList.Limit())
	if err != nil {
		return err
	}
	paginatedList.Items = items
	return c.Write(paginatedList)
}

func (r *cardResource) create(c *routing.Context) error {
	var model models.Card
	if err := c.Read(&model); err != nil {
		return err
	}
	response, err := r.service.Create(app.GetRequestScope(c), &model)
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

	model, err := r.service.Get(rs, id)
	if err != nil {
		return err
	}

	if err := c.Read(model); err != nil {
		return err
	}

	response, err := r.service.Update(rs, id, model)
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
