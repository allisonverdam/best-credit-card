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
		GetWallet(rs app.RequestScope, wallet_id int) (*models.Wallet, error)
		CreateWallet(rs app.RequestScope, wallet *models.Wallet) (*models.Wallet, error)
		UpdateWallet(rs app.RequestScope, wallet_id int, wallet *models.Wallet) (*models.Wallet, error)
		DeleteWallet(rs app.RequestScope, wallet_id int) (*models.Wallet, error)
		GetAuthenticatedPersonWallets(rs app.RequestScope) ([]models.Wallet, error)
	}

	// walletResource define os handlers para as chamadas do controller.
	walletResource struct {
		service walletService
	}
)

// ServeCard define as rotas.
func ServeWalletResource(rg *routing.RouteGroup, service walletService) {
	r := &walletResource{service}
	rg.Get("/wallets/<wallet_id>", r.GetWallet)
	rg.Get("/me/wallets", r.GetAuthenticatedPersonWallets)
	rg.Post("/wallets", r.CreateWallet)
	rg.Put("/wallets/<wallet_id>", r.UpdateWallet)
	rg.Delete("/wallets/<wallet_id>", r.DeleteWallet)
}

/**
* @api {get} /wallets/:wallet_id GetWallet
* @apiVersion 1.0.0
* @apiName GetWallet
* @apiDescription Retorna a carteira com o id passado por parametro.
* @apiGroup Wallet
* @apiUse AuthRequired
* @apiUse NotFoundError
*
* @apiError Forbidden A carteira não pertence ao usuário autenticado.
* @apiErrorExample Forbidden:
*     HTTP/1.1 403 Forbidden
*     {
*	"error_code": "FORBIDDEN",
*	"message": "You're not allowed to do this.",
*	"developer_message": "This wallet does not belong to the authenticated user."
*     }
* @apiSuccessExample Success-Response:
*     HTTP/1.1 200 OK
*      {
*	 "id": 1,
*	 "current_limit": 750,
*	 "maximum_limit": 1000,
*	 "avaliable_limit": 250,
*	 "person_id": 1
*      }
**/
func (r *walletResource) GetWallet(c *routing.Context) error {
	wallet_id, err := strconv.Atoi(c.Param("wallet_id"))
	if err != nil {
		return err
	}

	wallet, err := r.service.GetWallet(app.GetRequestScope(c), wallet_id)
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
func (r *walletResource) GetAuthenticatedPersonWallets(c *routing.Context) error {
	wallets, err := r.service.GetAuthenticatedPersonWallets(app.GetRequestScope(c))
	if err != nil {
		return err
	}

	return c.Write(wallets)
}

/**
* @api {post} /wallets CreateWallet
* @apiVersion 1.0.0
* @apiName CreateWallet
* @apiDescription Cria uma nova carteira.
* @apiGroup Wallet
* @apiUse AuthRequired
*
* @apiSuccessExample Success-Response:
*     HTTP/1.1 200 OK
*      {
*	 "id": 5,
*	 "current_limit": 0,
*	 "maximum_limit": 0,
*	 "avaliable_limit": 0,
*	 "person_id": 1
*      }
**/
func (r *walletResource) CreateWallet(c *routing.Context) error {
	tempWallet := models.Wallet{}
	tempWallet.PersonId = app.GetRequestScope(c).UserID()

	wallet, err := r.service.CreateWallet(app.GetRequestScope(c), &tempWallet)
	if err != nil {
		return err
	}

	return c.Write(wallet)
}

/**
* @api {put} /wallets/:wallet_id UpdateWallet
* @apiVersion 1.0.0
* @apiName UpdateWallet
* @apiDescription Atualizar uma carteira.
* @apiGroup Wallet
* @apiUse AuthRequired
* @apiUse NotFoundError
*
* @apiError Forbidden Essa carteira não pertence ao usuário autenticado.
* @apiErrorExample Forbidden:
*     HTTP/1.1 403 Forbidden
*     {
*	"error_code": "FORBIDDEN",
*	"message": "You're not allowed to do this.",
*	"developer_message": "This wallet does not belong to the authenticated user."
*     }
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
*	 "avaliable_limit": 0,
*	 "person_id": 1
*      }
**/
func (r *walletResource) UpdateWallet(c *routing.Context) error {
	wallet_id, err := strconv.Atoi(c.Param("wallet_id"))
	if err != nil {
		return err
	}

	rs := app.GetRequestScope(c)

	wallet, err := r.service.GetWallet(rs, wallet_id)
	if err != nil {
		return err
	}

	if err := c.Read(wallet); err != nil {
		return err
	}

	newWallet, err := r.service.UpdateWallet(rs, wallet_id, wallet)
	if err != nil {
		return err
	}

	return c.Write(newWallet)
}

/**
* @api {delete} /wallets/:wallet_id DeleteWallet
* @apiVersion 1.0.0
* @apiName DeleteWallet
* @apiDescription Apaga a carteira com o id passado por parametro.
* @apiGroup Wallet
* @apiUse AuthRequired
* @apiUse NotFoundError
*
* @apiError Forbidden A carteira não pertence ao usuário autenticado.
* @apiErrorExample Forbidden:
*     HTTP/1.1 403 Forbidden
*     {
*	"error_code": "FORBIDDEN",
*	"message": "You're not allowed to do this.",
*	"developer_message": "This wallet does not belong to the authenticated user."
*     }
* @apiSuccessExample Success-Response:
*     HTTP/1.1 200 OK
*      {
*	 "id": 1,
*	 "current_limit": 100,
*	 "maximum_limit": 200,
*	 "avaliable_limit": 100,
*	 "person_id": 1
*      }
**/
func (r *walletResource) DeleteWallet(c *routing.Context) error {
	wallet_id, err := strconv.Atoi(c.Param("wallet_id"))
	if err != nil {
		return err
	}

	wallet, err := r.service.DeleteWallet(app.GetRequestScope(c), wallet_id)
	if err != nil {
		return err
	}

	return c.Write(wallet)
}
