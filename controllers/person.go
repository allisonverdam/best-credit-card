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
		UpdateAuthenticatedPerson(rs app.RequestScope, person *models.Person) (*models.Person, error)
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
	rg.Put("/me", r.UpdateAuthenticatedPerson)
}

/**
* @api {get} /me GetPerson
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
* @api {put} /me UpdatePerson
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

	wallet, err := r.service.UpdateAuthenticatedPerson(rs, &person)
	if err != nil {
		return err
	}

	return c.Write(wallet)
}
