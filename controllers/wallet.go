package controllers

import (
	"strconv"

	"github.com/allisonverdam/best-credit-card/app"
	"github.com/allisonverdam/best-credit-card/models"
	routing "github.com/go-ozzo/ozzo-routing"
)

type (
	// walletService especifica a interface que é utilizada pelo walletResource
	walletService interface {
		Get(rs app.RequestScope, id int) (*models.Wallet, error)
		Create(rs app.RequestScope, model *models.Wallet) (*models.Wallet, error)
		Update(rs app.RequestScope, id int, model *models.Wallet) (*models.Wallet, error)
		Delete(rs app.RequestScope, id int) (*models.Wallet, error)
		// GetAuthenticatedPersonWallets(rs app.RequestScope, personId int) ([]models.Wallet, error)
	}

	// walletResource define os handlers para as chamadas do controller.
	walletResource struct {
		service walletService
	}
)

// ServeCard define as rotas.
func ServeWalletResource(rg *routing.RouteGroup, service walletService) {
	r := &walletResource{service}
	rg.Get("/wallets/<id>", r.Get)
	rg.Post("/wallets", r.Create)
	rg.Put("/wallets/<id>", r.Update)
	rg.Delete("/wallets/<id>", r.Delete)
}

func (r *walletResource) Get(c *routing.Context) error {
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

func (r *walletResource) Create(c *routing.Context) error {
	var model models.Wallet
	if err := c.Read(&model); err != nil {
		return err
	}
	response, err := r.service.Create(app.GetRequestScope(c), &model)
	if err != nil {
		return err
	}

	return c.Write(response)
}

func (r *walletResource) Update(c *routing.Context) error {
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

func (r *walletResource) Delete(c *routing.Context) error {
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
