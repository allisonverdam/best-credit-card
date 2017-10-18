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
		Register(rs app.RequestScope, model *models.Person) (*models.Person, error)
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

func (r *authResource) Register(c *routing.Context) error {
	var model models.Person
	if err := c.Read(&model); err != nil {
		return err
	}
	response, err := r.service.Register(app.GetRequestScope(c), &model)
	if err != nil {
		return err
	}

	return c.Write(response)
}

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
