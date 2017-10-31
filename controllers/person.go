package controllers

import (
	"github.com/allisonverdam/best-credit-card/app"
	"github.com/allisonverdam/best-credit-card/models"
	routing "github.com/go-ozzo/ozzo-routing"
)

type (
	// personService especifica a interface que é utilizada pelo personResource.
	personService interface {
		GetPerson(rs app.RequestScope, id int) (*models.Person, error)
		GetAuthenticatedPersonWallets(rs app.RequestScope) ([]models.Wallet, error)
		UpdateAuthenticatedPerson(rs app.RequestScope, id int, person *models.Person) (*models.Person, error)
	}

	// personResource define os handlers para as chamadas do controller.
	personResource struct {
		service personService
	}
)

// ServePersonResource define as rotas
func ServePersonResource(rg *routing.RouteGroup, service personService) {
	r := &personResource{service}
	rg.Get("/me", r.GetAuthenticatedPerson)
	rg.Get("/me/wallets", r.GetAuthenticatedPersonWallets)
	rg.Put("/me", r.UpdateAuthenticatedPerson)
}

/**
* @api {get} /me GetAuthenticatedPerson
* @apiVersion 1.0.0
* @apiName GetAuthenticatedPerson
* @apiDescription Retorna o usuário autenticado.
* @apiGroup Person
* @apiUse AuthRequired
*
* @apiSuccessExample Success-Response:
*     HTTP/1.1 200 OK
*      {
*	 "id": 1,
*	 "name": "Allison V.",
*	 "username": "allisonverdam",
*	 "email": "allison@g.com"
*      }
**/
func (r *personResource) GetAuthenticatedPerson(c *routing.Context) error {
	wallet, err := r.service.GetPerson(app.GetRequestScope(c), app.GetRequestScope(c).UserID())
	if err != nil {
		return err
	}

	return c.Write(wallet)
}

/**
* @api {get} /me/wallets GetAuthenticatedPersonWallets
* @apiVersion 1.0.0
* @apiName GetAuthenticatedPersonWallets
* @apiDescription Retorna as carteiras do usuário autenticado.
* @apiGroup Person
* @apiUse AuthRequired
*
* @apiSuccessExample Success-Response:
*     HTTP/1.1 200 OK
*     [
*      {
*	 "id": 3,
*	 "id": 1,
*	 "real_limit": 0,
*	 "maximum_limit": 0,
*	 "person_id": 1
*      },
*      {
*	 "id": 4,
*	 "real_limit": 0,
*	 "maximum_limit": 0,
*	 "person_id": 1
*      }
**/
func (r *personResource) GetAuthenticatedPersonWallets(c *routing.Context) error {
	wallets, err := r.service.GetAuthenticatedPersonWallets(app.GetRequestScope(c))
	if err != nil {
		return err
	}

	return c.Write(wallets)
}

/**
* @api {put} /me UpdateAuthenticatedPerson
* @apiVersion 1.0.0
* @apiName UpdateAuthenticatedPerson
* @apiDescription Atualiza o usuário autenticado.
* @apiGroup Person
* @apiUse AuthRequired
*
* @apiParamExample {json} Request-Example:
*     {
*       "email": "allison2222@g.com",
*       "name": "allison",
*       "username": "allisonverdam"
*     }
* @apiSuccessExample Success-Response:
*     HTTP/1.1 200 OK
*      {
*	 "id": 1,
*	 "name": "Allison V.",
*	 "username": "allisonverdam",
*	 "email": "allison2222@g.com"
*      }
**/
func (r *personResource) UpdateAuthenticatedPerson(c *routing.Context) error {
	rs := app.GetRequestScope(c)

	person := models.Person{}

	if err := c.Read(&person); err != nil {
		return err
	}

	wallet, err := r.service.UpdateAuthenticatedPerson(rs, rs.UserID(), &person)
	if err != nil {
		return err
	}

	return c.Write(wallet)
}
