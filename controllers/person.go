package controllers

import (
	"strconv"

	"github.com/allisonverdam/go-api-mcc/app"
	"github.com/allisonverdam/go-api-mcc/models"
	routing "github.com/go-ozzo/ozzo-routing"
)

type (
	// personService especifica a interface que é utilizada pelo personResource.
	personService interface {
		Get(rs app.RequestScope, id int) (*models.Person, error)
		GetPersonByName(rs app.RequestScope, personname string) (*models.Person, error)
		Query(rs app.RequestScope, offset, limit int) ([]models.Person, error)
		Count(rs app.RequestScope) (int, error)
		Create(rs app.RequestScope, model *models.Person) (*models.Person, error)
		Update(rs app.RequestScope, id int, model *models.Person) (*models.Person, error)
		Delete(rs app.RequestScope, id int) (*models.Person, error)
	}

	// personResource define os handlers para as chamadas do controller.
	personResource struct {
		service personService
	}
)

// ServePersonResource define as rotas
func ServePersonResource(rg *routing.RouteGroup, service personService) {
	r := &personResource{service}
	rg.Get("/persons/<id>", r.get)
	rg.Get("/persons/<personname>", r.GetPersonByName)
	rg.Get("/persons", r.query)
	rg.Post("/persons", r.create)
	rg.Put("/persons/<id>", r.update)
	rg.Delete("/persons/<id>", r.delete)
}

func (r *personResource) get(c *routing.Context) error {
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

func (r *personResource) GetPersonByName(c *routing.Context) error {
	response, err := r.service.GetPersonByName(app.GetRequestScope(c), c.Param("personname"))
	if err != nil {
		return err
	}

	return c.Write(response)
}

func (r *personResource) query(c *routing.Context) error {
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

func (r *personResource) create(c *routing.Context) error {
	var model models.Person
	if err := c.Read(&model); err != nil {
		return err
	}
	response, err := r.service.Create(app.GetRequestScope(c), &model)
	if err != nil {
		return err
	}

	return c.Write(response)
}

func (r *personResource) update(c *routing.Context) error {
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

func (r *personResource) delete(c *routing.Context) error {
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
