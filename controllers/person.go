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
		GetAuthenticatedPersonWallets(rs app.RequestScope, personId int) ([]models.Wallet, error)
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
	rg.Get("/me", r.Get)
	rg.Get("/me/wallets", r.GetAuthenticatedPersonWallets)
	rg.Put("/me", r.Update)
}

func (r *personResource) Get(c *routing.Context) error {
	response, err := r.service.Get(app.GetRequestScope(c), app.GetRequestScope(c).UserID())
	if err != nil {
		return err
	}

	return c.Write(response)
}

func (r *personResource) GetAuthenticatedPersonWallets(c *routing.Context) error {
	response, err := r.service.GetAuthenticatedPersonWallets(app.GetRequestScope(c), app.GetRequestScope(c).UserID())
	if err != nil {
		return err
	}

	return c.Write(response)
}

func (r *personResource) Update(c *routing.Context) error {
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
