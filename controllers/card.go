package controllers

import (
	"strconv"

	"github.com/allisonverdam/best-credit-card/app"
	"github.com/allisonverdam/best-credit-card/models"
	routing "github.com/go-ozzo/ozzo-routing"
)

type (
	// cardService especifica a interface que é utilizada pelo cardResource
	cardService interface {
		GetCard(rs app.RequestScope, card_id int) (*models.Card, error)
		GetBestCards(rs app.RequestScope, order *models.Order) ([]models.Card, error)
		GetAuthenticatedPersonCards(rs app.RequestScope) ([]models.Card, error)
		PayCreditCard(rs app.RequestScope, order models.Order) (*models.Card, error)
		CreateCard(rs app.RequestScope, card *models.Card) (*models.Card, error)
		UpdateCard(rs app.RequestScope, card_id int, card *models.Card) (*models.Card, error)
		DeleteCard(rs app.RequestScope, card_id int) (*models.Card, error)
	}

	// cardResource define os handlers para as chamadas do controller.
	cardResource struct {
		service cardService
	}
)

// ServeCard define as rotas.
func ServeCardResource(rg *routing.RouteGroup, service cardService) {
	r := &cardResource{service}
	rg.Get("/cards/<card_id>", r.GetCard)
	rg.Get("/cards", r.GetCards)
	rg.Post("/cards", r.CreateCard)
	rg.Post("/cards/pay", r.PayCreditCard)
	rg.Post("/cards/best-card", r.GetBestCards)
	rg.Put("/cards/<card_id>", r.UpdateCard)
	rg.Delete("/cards/<card_id>", r.DeleteCard)
}

/**
* @api {post} /cards/best-card GetBestCards
* @apiVersion 1.0.0
* @apiName GetBestCards
* @apiGroup Card
* @apiDescription Retorna o melhor cartão para a compra.
*
* @apiUse AuthRequired
*
* @apiParamExample {json} Request-Example:
*     {
*       "price": 100
*     }
*
*
* @apiSuccessExample Success-Response:
*     HTTP/1.1 200 OK
*     [
*      {
*	 "id": 3,
*	 "number": "1234123412341232",
*	 "due_date": 11,
*	 "expiration_month": 8,
*	 "expiration_year": 16,
*	 "currency": "BRL",
*	 "cvv": 123,
*	 "real_limit": 500,
*	 "avaliable_limit": 450
*      }
*     ]
*
* @apiUse NotFoundError
* @apiUse ValidatePrice
*
* @apiError ValidatePrice O parametro 'price' não pode ser menor que 0.
 */
func (r *cardResource) GetBestCards(c *routing.Context) error {
	var order models.Order
	if err := c.Read(&order); err != nil {
		return err
	}
	if err := order.ValidatePrice(); err != nil {
		return err
	}

	card, err := r.service.GetBestCards(app.GetRequestScope(c), &order)
	if err != nil {
		return err
	}

	return c.Write(card)
}

/**
* @api {post} /cards/pay PayCreditCard
* @apiVersion 1.0.0
* @apiName PayCreditCard
* @apiDescription Pagar um cartão para liberar crédito.
* @apiGroup Card
* @apiUse AuthRequired
* @apiUse NotFoundError
* @apiUse ValidatePrice
*
* @apiError Forbidden O parametro 'wallet_id' informado não pertence ao usuário autenticado.
* @apiError ValidatePrice O parametro 'price' não pode ser menor que 0.
* @apiErrorExample Forbidden:
*     HTTP/1.1 403 Forbidden
*     {
*	"error_code": "FORBIDDEN",
*	"message": "You're not allowed to do this.",
*	"developer_message": "This card does not belong to the authenticated user."
*     }
* @apiParamExample {json} Request-Example:
*     {
*       "price": 100,
		"card_id": 1
*     }
*
* @apiSuccessExample Success-Response:
*     HTTP/1.1 200 OK
*      {
*	 "id": 1,
*	 "number": "1234123412341232",
*	 "due_date": 11,
*	 "expiration_month": 8,
*	 "expiration_year": 16,
*	 "currency": "BRL",
*	 "cvv": 123,
*	 "real_limit": 500,
*	 "avaliable_limit": 450
*      }
**/
func (r *cardResource) PayCreditCard(c *routing.Context) error {
	var order models.Order
	if err := c.Read(&order); err != nil {
		return err
	}
	if err := order.ValidateCardIdAndPrice(); err != nil {
		return err
	}

	card, err := r.service.PayCreditCard(app.GetRequestScope(c), *&order)
	if err != nil {
		return err
	}

	return c.Write(card)
}

/**
* @api {GetCard} /cards/:card_id GetCard
* @apiVersion 1.0.0
* @apiName GetCard
* @apiDescription Retorna o cartao com o id passado por parametro.
* @apiGroup Card
* @apiUse AuthRequired
* @apiUse NotFoundError
*
* @apiError Forbidden O cartão não pertence ao usuário autenticado.
* @apiErrorExample Forbidden:
*     HTTP/1.1 403 Forbidden
*     {
*	"error_code": "FORBIDDEN",
*	"message": "You're not allowed to do this.",
*	"developer_message": "This card does not belong to the authenticated user."
*     }
* @apiSuccessExample Success-Response:
*     HTTP/1.1 200 OK
*      {
*	 "id": 3,
*	 "number": "1234123412341232",
*	 "due_date": 11,
*	 "expiration_month": 8,
*	 "expiration_year": 16,
*	 "currency": "BRL",
*	 "cvv": 123,
*	 "real_limit": 500,
*	 "avaliable_limit": 450,
*	 "wallet_id": 1
*      }
**/
func (r *cardResource) GetCard(c *routing.Context) error {
	card_id, err := strconv.Atoi(c.Param("card_id"))
	if err != nil {
		return err
	}

	card, err := r.service.GetCard(app.GetRequestScope(c), card_id)
	if err != nil {
		return err
	}

	return c.Write(card)
}

/**
* @api {get} /cards GetCards
* @apiVersion 1.0.0
* @apiName GetCards
* @apiDescription Retorna todos os cartões da carteira do usuario autenticado.
* @apiGroup Card
* @apiUse AuthRequired
*
* @apiSuccessExample Success-Response:
*     HTTP/1.1 200 OK
*     [
*      {
*	 "id": 3,
*	 "number": "1234123412341232",
*	 "due_date": 11,
*	 "expiration_month": 8,
*	 "expiration_year": 16,
*	 "currency": "BRL",
*	 "cvv": 123,
*	 "real_limit": 500,
*	 "avaliable_limit": 450
*      }
*     ]
**/
func (r *cardResource) GetCards(c *routing.Context) error {
	cards, err := r.service.GetAuthenticatedPersonCards(app.GetRequestScope(c))
	if err != nil {
		return err
	}

	return c.Write(cards)
}

/**
* @api {post} /cards CreateCard
* @apiVersion 1.0.0
* @apiName CreateCard
* @apiDescription Cria um novo cartão.
* @apiGroup Card
* @apiUse AuthRequired
* @apiUse NotFoundError
*
* @apiParamExample {json} Request-Example:
*      {
*	 "number": "1234123412341232",
*	 "due_date": 11,
*	 "expiration_month": 8,
*	 "expiration_year": 16,
*	 "currency": "BRL",
*	 "cvv": 123,
*	 "real_limit": 500,
*	 "avaliable_limit": 450
*      }
*
* @apiSuccessExample Success-Response:
*     HTTP/1.1 200 OK
*      {
*	 "id": 3,
*	 "number": "1234123412341232",
*	 "due_date": 11,
*	 "expiration_month": 8,
*	 "expiration_year": 16,
*	 "currency": "BRL",
*	 "cvv": 123,
*	 "real_limit": 500,
*	 "avaliable_limit": 450
*      }
**/
func (r *cardResource) CreateCard(c *routing.Context) error {
	card := models.Card{}
	if err := c.Read(&card); err != nil {
		return err
	}

	if err := card.Validate(); err != nil {
		return err
	}

	cardBD, err := r.service.CreateCard(app.GetRequestScope(c), &card)
	if err != nil {
		return err
	}

	return c.Write(cardBD)
}

/**
* @api {put} /cards/:card_id UpdateCard
* @apiVersion 1.0.0
* @apiName UpdateCard
* @apiDescription Atualizar um cartão.
* @apiGroup Card
* @apiUse AuthRequired
* @apiUse NotFoundError
*
* @apiError Forbidden Este cartão não pertence ao usuário autenticado.
* @apiErrorExample Forbidden:
*     HTTP/1.1 403 Forbidden
*     {
*	"error_code": "FORBIDDEN",
*	"message": "You're not allowed to do this.",
*	"developer_message": "This card does not belong to the authenticated user."
*     }
* @apiParamExample {json} Request-Example:
*      {
*	 "number": "1234123412341232",
*	 "due_date": 11,
*	 "expiration_month": 8,
*	 "expiration_year": 16,
*	 "currency": "BRL",
*	 "cvv": 123,
*	 "real_limit": 700,
*	 "avaliable_limit": 550
*      }
*
* @apiSuccessExample Success-Response:
*     HTTP/1.1 200 OK
*      {
*	 "id": 3,
*	 "number": "1234123412341232",
*	 "due_date": 11,
*	 "expiration_month": 8,
*	 "expiration_year": 16,
*	 "currency": "BRL",
*	 "cvv": 123,
*	 "real_limit": 700,
*	 "avaliable_limit": 550
*      }
**/
func (r *cardResource) UpdateCard(c *routing.Context) error {
	card_id, err := strconv.Atoi(c.Param("card_id"))
	if err != nil {
		return err
	}

	card := models.Card{}
	if err := c.Read(&card); err != nil {
		return err
	}

	cardBD, err := r.service.UpdateCard(app.GetRequestScope(c), card_id, &card)
	if err != nil {
		return err
	}

	return c.Write(cardBD)
}

/**
* @api {DeleteCard} /cards/:card_id DeleteCard
* @apiVersion 1.0.0
* @apiName DeleteCard
* @apiDescription Apaga o cartao com o id passado por parametro.
* @apiGroup Card
* @apiUse AuthRequired
* @apiUse NotFoundError
*
* @apiError Forbidden Esse cartão não pertence ao usuário autenticado.
* @apiErrorExample Forbidden:
*     HTTP/1.1 403 Forbidden
*     {
*	"error_code": "FORBIDDEN",
*	"message": "You're not allowed to do this.",
*	"developer_message": "This card does not belong to the authenticated user."
*     }
* @apiSuccessExample Success-Response:
*     HTTP/1.1 200 OK
*      {
*	 "id": 3,
*	 "number": "1234123412341232",
*	 "due_date": 11,
*	 "expiration_month": 8,
*	 "expiration_year": 16,
*	 "currency": "BRL",
*	 "cvv": 123,
*	 "real_limit": 500,
*	 "avaliable_limit": 450
*      }
**/
func (r *cardResource) DeleteCard(c *routing.Context) error {
	card_id, err := strconv.Atoi(c.Param("card_id"))
	if err != nil {
		return err
	}

	card, err := r.service.DeleteCard(app.GetRequestScope(c), card_id)
	if err != nil {
		return err
	}

	return c.Write(card)
}
