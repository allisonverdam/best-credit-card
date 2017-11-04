package controllers

import (
	"github.com/allisonverdam/best-credit-card/app"
	"github.com/allisonverdam/best-credit-card/errors"
	"github.com/allisonverdam/best-credit-card/models"
	routing "github.com/go-ozzo/ozzo-routing"
)

type (
	// authService especifica a interface que é utilizada pelo cardResource
	authService interface {
		Register(rs app.RequestScope, person *models.Person) (*models.Person, error)
		Login(c *routing.Context, credential models.Credential, signingKey string) (string, error)
	}

	// authResource define os handlers para as chamadas do controller.
	authResource struct {
		service authService
	}
)

// ServeAuthResource define as rotas.
func ServeAuthResource(rg *routing.RouteGroup, service authService) {
	r := &authResource{service}
	rg.Post("/register", r.Register)
	rg.Post("/login", r.Login)
}

/**
* @api {post} /register Register
* @apiVersion 1.0.0
* @apiName Register
* @apiGroup Auth
* @apiDescription Registra um novo usuário.
*
* @apiUse ContentTypeJson
*
* @apiParamExample {json} Request-Example:
*     {
*       "email":"amanda@gmail.com",
*       "name":"amanda",
*       "password":"as232ff",
*       "username":"amanda"
*     }
*
*
* @apiSuccessExample Success-Response:
*     HTTP/1.1 200 OK
*      {
*	 "id": 4,
*	 "name": "amanda",
*	 "username": "amanda",
*	 "email": "amanda@gmail.com"
*      }
*
 */
func (r *authResource) Register(c *routing.Context) error {
	person := models.Person{}
	if err := c.Read(&person); err != nil {
		return err
	}

	response, err := r.service.Register(app.GetRequestScope(c), &person)
	if err != nil {
		return err
	}

	return c.Write(response)
}

/**
* @apiDefine InvalidCredentials
* @apiError InvalidCredentials Login ou senha inválido.
* @apiErrorExample InvalidCredentials:
*     HTTP/1.1 401 Unauthorized
*     {
*	"error_code": "UNAUTHORIZED",
*	"message": "Authentication failed.",
*	"developer_message": "Authentication failed: Authentication failed."
*     }
 */

/**
* @api {post} /login Login
* @apiVersion 1.0.0
* @apiName Login
* @apiGroup Auth
* @apiDescription Autentica o usuário.
*
* @apiUse ContentTypeJson
* @apiUse InvalidCredentials
*
* @apiParamExample {json} Request-Example:
*     {
*       "password":"as232ff",
*       "username":"amanda"
*     }
*
*
* @apiSuccessExample Success-Response:
*     HTTP/1.1 200 OK
*      {
*	 "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MDk0OTI0NjIsImlkIjoxLCJuYW1lIjoiQWxsaXNvbiBWLiJ9.hRpe6GDdZVqGYVNAl8OfPdoqyWfJRfwRG1i3PsM_ay0"
*      }
*
 */
func (r *authResource) Login(c *routing.Context) error {
	var credential models.Credential
	if err := c.Read(&credential); err != nil {
		return errors.Unauthorized(err.Error())
	}

	token, err := r.service.Login(c, credential, app.Config.JWTSigningKey)

	if err != nil {
		return errors.Unauthorized(err.Error())
	}

	return c.Write(map[string]string{
		"token": token,
	})
}
