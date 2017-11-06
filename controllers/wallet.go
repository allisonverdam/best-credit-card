package controllers

import (
	"github.com/allisonverdam/best-credit-card/app"
	"github.com/allisonverdam/best-credit-card/models"
	routing "github.com/go-ozzo/ozzo-routing"
)

type (
	// walletService especifica a interface que é utilizada pelo walletResource
	walletService interface {
		GetWallet(rs app.RequestScope, wallet_id int) (*models.Wallet, error)
		CreateWallet(rs app.RequestScope, wallet *models.Wallet) (*models.Wallet, error)
		UpdateAuthenticatedPersonWallet(rs app.RequestScope, wallet *models.Wallet) (*models.Wallet, error)
		GetAuthenticatedPersonWallet(rs app.RequestScope) (models.Wallet, error)
	}

	// walletResource define os handlers para as chamadas do controller.
	walletResource struct {
		service walletService
	}
)

// ServeCard define as rotas.
func ServeWalletResource(rg *routing.RouteGroup, service walletService) {
	r := &walletResource{service}
	rg.Get("/wallets", r.GetAuthenticatedPersonWallet)
	rg.Put("/wallets", r.UpdateAuthenticatedPersonWallet)
}

/**
* @api {get} /wallets GetWallet
* @apiVersion 1.0.0
* @apiName GetAuthenticatedPersonWallet
* @apiDescription Retorna a carteira do usuário autenticado.
* @apiGroup Wallet
* @apiUse AuthRequired
*
* @apiSuccessExample Success-Response:
*     HTTP/1.1 200 OK
*     {
*	 "current_limit": 700,
*	 "maximum_limit": 1000
*	 "avaliable_limit": 0,
*     }
**/
func (r *walletResource) GetAuthenticatedPersonWallet(c *routing.Context) error {
	wallets, err := r.service.GetAuthenticatedPersonWallet(app.GetRequestScope(c))
	if err != nil {
		return err
	}

	return c.Write(wallets)
}

/**
* @api {put} /wallets UpdateWallet
* @apiVersion 1.0.0
* @apiName UpdateAuthenticatedPersonWallet
* @apiDescription Atualiza a carteira do usuário autenticado.
* @apiGroup Wallet
* @apiUse AuthRequired
* @apiUse NotFoundError
*
* @apiParamExample {json} Request-Example:
*      {
*	 "current_limit": 700
*      }
*
* @apiSuccessExample Success-Response:
*     HTTP/1.1 200 OK
*      {
*	 "id": 5,
*	 "current_limit": 700,
*	 "maximum_limit": 2000,
*	 "avaliable_limit": 0
*      }
**/
func (r *walletResource) UpdateAuthenticatedPersonWallet(c *routing.Context) error {
	rs := app.GetRequestScope(c)

	wallet := models.Wallet{}
	if err := c.Read(&wallet); err != nil {
		return err
	}

	newWallet, err := r.service.UpdateAuthenticatedPersonWallet(rs, &wallet)
	if err != nil {
		return err
	}

	return c.Write(newWallet)
}
