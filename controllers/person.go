package controllers

import (
	"github.com/allisonverdam/best-credit-card/app"
	"github.com/allisonverdam/best-credit-card/models"
	routing "github.com/go-ozzo/ozzo-routing"
)

type (
	// personService especifica a interface que é utilizada pelo personResource.
	personService interface {
		Get(rs app.RequestScope, id int) (*models.Person, error)
		Update(rs app.RequestScope, id int, person *models.Person) (*models.Person, error)
	}

	// personResource define os handlers para as chamadas do controller.
	personResource struct {
		service personService
	}
)

// ServePersonResource define as rotas
func ServePersonResource(rg *routing.RouteGroup, service personService) {
	r := &personResource{service}
	rg.Get("/me", r.get)
	rg.Put("/me", r.update)
}

func (r *personResource) get(c *routing.Context) error {
	response, err := r.service.Get(app.GetRequestScope(c), app.GetRequestScope(c).UserID())
	if err != nil {
		return err
	}

	return c.Write(response)
}

func (r *personResource) update(c *routing.Context) error {
	rs := app.GetRequestScope(c)

	person := models.Person{}

	if err := c.Read(&person); err != nil {
		return err
	}

	response, err := r.service.Update(rs, rs.UserID(), &person)
	if err != nil {
		return err
	}

	return c.Write(response)
}
