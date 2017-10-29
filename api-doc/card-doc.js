/**
* @api {post} /cards/best-card GetBestCards - Retorna o melhor cartão para a compra.
* @apiVersion 1.0.0
* @apiName GetBestCards
* @apiGroup Card
*
* @apiUse AuthRequired
*
* @apiParamExample {json} Request-Example:
*     {
*       "price": 100,
		"wallet_id": 1
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
*	 "cvv": 123,
*	 "real_limit": 500,
*	 "current_limit": 450,
*	 "wallet_id": 1
*      }
*     ]
*
* @apiUse UserNotFoundError
* @apiUse ValidatePrice
*
* @apiError Forbidden O parametro 'wallet_id' informado não pertence ao usuário autenticado.
* @apiError ValidatePrice O parametro 'price' não pode ser menor que 0.
* @apiErrorExample Forbidden:
*     HTTP/1.1 403 Forbidden
*     {
*	"error_code": "FORBIDDEN",
*	"message": "This wallet does not belong to this user."
*     }
*/

/**
* @api {post} /cards/pay PayCreditCard - Pagar um cartão para liberar crédito.
* @apiVersion 1.0.0
* @apiName PayCreditCard
* @apiGroup Card
* @apiUse AuthRequired
* @apiUse UserNotFoundError
* @apiUse ValidatePrice
*
* @apiError Forbidden O parametro 'wallet_id' informado não pertence ao usuário autenticado.
* @apiError ValidatePrice O parametro 'price' não pode ser menor que 0.
* @apiErrorExample Forbidden:
*     HTTP/1.1 403 Forbidden
*     {
*	"error_code": "FORBIDDEN",
*	"message": "This card does not belong to the authenticated user."
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
*	 "cvv": 123,
*	 "real_limit": 500,
*	 "current_limit": 450,
*	 "wallet_id": 1
*      }
**/

/**
* @api {get} /cards/:card_id GetCard - Retorna o cartao com o id passado por parametro.
* @apiVersion 1.0.0
* @apiName GetCard
* @apiGroup Card
* @apiUse AuthRequired
* @apiUse UserNotFoundError
*
* @apiError Forbidden O cartão não pertence ao usuário autenticado.
* @apiErrorExample Forbidden:
*     HTTP/1.1 403 Forbidden
*     {
*	"error_code": "FORBIDDEN",
*	"message": "This card does not belong to the authenticated user."
*     }
* @apiSuccessExample Success-Response:
*     HTTP/1.1 200 OK
*      {
*	 "id": 3,
*	 "number": "1234123412341232",
*	 "due_date": 11,
*	 "expiration_month": 8,
*	 "expiration_year": 16,
*	 "cvv": 123,
*	 "real_limit": 500,
*	 "current_limit": 450,
*	 "wallet_id": 1
*      }
**/

/**
* @api {get} /cards/:wallet_id CardsWallet - Retorna a lista de cartões de uma determinada carteira.
* @apiVersion 1.0.0
* @apiName CardsWallet
* @apiGroup Card
* @apiUse AuthRequired
* @apiUse UserNotFoundError
*
* @apiError Forbidden O cartão não pertence ao usuário autenticado.
* @apiErrorExample Forbidden:
*     HTTP/1.1 403 Forbidden
*     {
*	"error_code": "FORBIDDEN",
*	"message": "This wallet does not belong to the authenticated user."
*     }
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
*	 "cvv": 123,
*	 "real_limit": 500,
*	 "current_limit": 450,
*	 "wallet_id": 1
*      }
*     ]
**/

/**
* @api {post} /cards Create - Cria um novo cartão.
* @apiVersion 1.0.0
* @apiName Create
* @apiGroup Card
* @apiUse AuthRequired
* @apiUse UserNotFoundError
*
* @apiError Forbidden Essa carteira não pertence ao usuário autenticado.
* @apiErrorExample Forbidden:
*     HTTP/1.1 403 Forbidden
*     {
*	"error_code": "FORBIDDEN",
*	"message": "This wallet does not belong to the authenticated user."
*     }
* @apiParamExample {json} Request-Example:
*      {
*	 "number": "1234123412341232",
*	 "due_date": 11,
*	 "expiration_month": 8,
*	 "expiration_year": 16,
*	 "cvv": 123,
*	 "real_limit": 500,
*	 "current_limit": 450,
*	 "wallet_id": 1
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
*	 "cvv": 123,
*	 "real_limit": 500,
*	 "current_limit": 450,
*	 "wallet_id": 1
*      }
**/

/**
* @api {delete} /cards/:card_id Delete - Apaga o cartao com o id passado por parametro.
* @apiVersion 1.0.0
* @apiName Delete
* @apiGroup Card
* @apiUse AuthRequired
* @apiUse UserNotFoundError
*
* @apiError Forbidden O cartão não pertence ao usuário autenticado.
* @apiErrorExample Forbidden:
*     HTTP/1.1 403 Forbidden
*     {
*	"error_code": "FORBIDDEN",
*	"message": "This card does not belong to the authenticated user."
*     }
* @apiSuccessExample Success-Response:
*     HTTP/1.1 200 OK
*      {
*	 "id": 3,
*	 "number": "1234123412341232",
*	 "due_date": 11,
*	 "expiration_month": 8,
*	 "expiration_year": 16,
*	 "cvv": 123,
*	 "real_limit": 500,
*	 "current_limit": 450,
*	 "wallet_id": 1
*      }
**/

/**
* @api {put} /cards/:card_id Update - Atualizar um cartão.
* @apiVersion 1.0.0
* @apiName Update
* @apiGroup Card
* @apiUse AuthRequired
* @apiUse UserNotFoundError
*
* @apiError Forbidden Essa carteira não pertence ao usuário autenticado.
* @apiErrorExample Forbidden:
*     HTTP/1.1 403 Forbidden
*     {
*	"error_code": "FORBIDDEN",
*	"message": "This card does not belong to the authenticated user."
*     }
* @apiParamExample {json} Request-Example:
*      {
*	 "real_limit": 700,
*	 "current_limit": 550,
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
*	 "cvv": 123,
*	 "real_limit": 700,
*	 "current_limit": 550,
*	 "wallet_id": 1
*      }
**/