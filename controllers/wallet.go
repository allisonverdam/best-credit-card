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
		Get(rs app.RequestScope, wallet_id int) (*models.Wallet, error)
		Create(rs app.RequestScope, wallet *models.Wallet) (*models.Wallet, error)
		Update(rs app.RequestScope, wallet_id int, wallet *models.Wallet) (*models.Wallet, error)
		Delete(rs app.RequestScope, wallet_id int) (*models.Wallet, error)
	}

	// walletResource define os handlers para as chamadas do controller.
	walletResource struct {
		service walletService
	}
)

// ServeCard define as rotas.
func ServeWalletResource(rg *routing.RouteGroup, service walletService) {
	r := &walletResource{service}
	rg.Get("/wallets/<wallet_id>", r.Get)
	rg.Post("/wallets", r.Create)
	rg.Put("/wallets/<wallet_id>", r.Update)
	rg.Delete("/wallets/<wallet_id>", r.Delete)
}

func (r *walletResource) Get(c *routing.Context) error {
	wallet_id, err := strconv.Atoi(c.Param("wallet_id"))
	if err != nil {
		return err
	}

	wallet, err := r.service.Get(app.GetRequestScope(c), wallet_id)
	if err != nil {
		return err
	}

	return c.Write(wallet)
}

func (r *walletResource) Create(c *routing.Context) error {
	tempWallet := models.Wallet{}
	tempWallet.PersonId = app.GetRequestScope(c).UserID()

	wallet, err := r.service.Create(app.GetRequestScope(c), &tempWallet)
	if err != nil {
		return err
	}

	return c.Write(wallet)
}

func (r *walletResource) Update(c *routing.Context) error {
	wallet_id, err := strconv.Atoi(c.Param("wallet_id"))
	if err != nil {
		return err
	}

	rs := app.GetRequestScope(c)

	wallet, err := r.service.Get(rs, wallet_id)
	if err != nil {
		return err
	}

	if err := c.Read(wallet); err != nil {
		return err
	}

	newWallet, err := r.service.Update(rs, wallet_id, wallet)
	if err != nil {
		return err
	}

	return c.Write(newWallet)
}

func (r *walletResource) Delete(c *routing.Context) error {
	wallet_id, err := strconv.Atoi(c.Param("wallet_id"))
	if err != nil {
		return err
	}

	wallet, err := r.service.Delete(app.GetRequestScope(c), wallet_id)
	if err != nil {
		return err
	}

	return c.Write(wallet)
}
